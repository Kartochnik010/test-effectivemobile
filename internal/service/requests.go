package service

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	"github.com/rs/zerolog/log"
)

type requestsService struct {
	ErrorsChan chan error
}

func NewRequestsService() requests {
	return &requestsService{}
}

type ConcurrentlySafePerson struct {
	wg *sync.WaitGroup
	m  *sync.Mutex
	p  *models.Person
}

func (r *requestsService) AddData(person *models.Person) {
	urls := map[string]string{
		"age":         fmt.Sprintf("https://api.agify.io/?name=%s", person.Name),
		"gender":      fmt.Sprintf("https://api.genderize.io/?name=%s", person.Name),
		"nationality": fmt.Sprintf("https://api.nationalize.io/?name=%s", person.Name),
	}
	p := ConcurrentlySafePerson{p: person, m: &sync.Mutex{}, wg: &sync.WaitGroup{}}
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
		p.p.Age = sql.NullInt32{Int32: int32(age), Valid: true}
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
		p.p.Gender = sql.NullString{String: gender, Valid: true}
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
		p.p.Nationality = sql.NullString{String: countryID, Valid: true}
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
