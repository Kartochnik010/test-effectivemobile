package service

import "encoding/json"

type AgeResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type GenderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type CountryResponse struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func ParseAgeResponse(jsonData []byte) (int, error) {
	var ageResponse AgeResponse
	err := json.Unmarshal(jsonData, &ageResponse)
	if err != nil {
		return 0, err
	}
	return ageResponse.Age, nil
}

func ParseGenderResponse(jsonData []byte) (string, error) {
	var genderResponse GenderResponse
	err := json.Unmarshal(jsonData, &genderResponse)
	if err != nil {
		return "", err
	}
	return genderResponse.Gender, nil
}

func ParseCountryResponse(jsonData []byte) (string, error) {
	var countryResponse CountryResponse
	err := json.Unmarshal(jsonData, &countryResponse)
	if err != nil {
		return "", err
	}
	if len(countryResponse.Country) > 0 {
		return countryResponse.Country[0].CountryID, nil
	}
	return "", nil
}
