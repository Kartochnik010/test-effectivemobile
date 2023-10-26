package service

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	"github.com/rs/zerolog/log"
)

type RequestsService struct {
	ErrorsChan chan error
}

func NewRequestsService() Requests {
	return &RequestsService{}
}

type ConcurrentlySafePerson struct {
	wg *sync.WaitGroup
	m  *sync.Mutex
	p  *models.Person
}

func (r *RequestsService) AddData(person *models.Person) {
	urls := map[string]string{
		"age":         fmt.Sprintf("https://api.agify.io/?name=%s", person.Name),
		"gender":      fmt.Sprintf("https://api.genderize.io/?name=%s", person.Name),
		"nationality": fmt.Sprintf("https://api.nationalize.io/?name=%s", person.Name),
	}
	p := ConcurrentlySafePerson{p: person}
	p.wg.Add(3)
	go func(url string, p ConcurrentlySafePerson) {
		content, err := getContent(url)
		if err != nil {
			r.ErrorsChan <- err
			log.Error().Msg(err.Error())
		}
		age, err := ParseAgeResponse(content)
		if err != nil {
			log.Error().Msg(err.Error())
			r.ErrorsChan <- err
		}
		p.m.Lock()
		p.p.Age = age
		p.m.Unlock()
		p.wg.Done()
	}(urls["age"], p)
	go func(url string, p ConcurrentlySafePerson) {
		content, err := getContent(url)
		if err != nil {
			r.ErrorsChan <- err
			log.Error().Msg(err.Error())
		}
		gender, err := ParseGenderResponse(content)
		if err != nil {
			log.Error().Msg(err.Error())
			r.ErrorsChan <- err
		}
		p.m.Lock()
		p.p.Gender = gender
		p.m.Unlock()
		p.wg.Done()
	}(urls["gender"], p)
	go func(url string, p ConcurrentlySafePerson) {
		content, err := getContent(url)
		if err != nil {
			r.ErrorsChan <- err
			log.Error().Msg(err.Error())
		}
		countryID, err := ParseCountryResponse(content)
		if err != nil {
			log.Error().Msg(err.Error())
			r.ErrorsChan <- err
		}
		p.m.Lock()
		p.p.Nationality = countryID
		p.m.Unlock()
		p.wg.Done()
	}(urls["nationality"], p)
	p.wg.Wait()
}

func getContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
