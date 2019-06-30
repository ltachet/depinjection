package meteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Interface contains necessary functions to handle meteo.
type Interface interface {
	GetBestCityMeteo(cities []string) (string, error)
}

// Service contains info to handle meteo info.
type Service struct {
	client     *http.Client
	serverAddr string
}

// Meteo contains the meteo client response.
type Meteo struct {
	Data  string `json:"data"`
	Grade int    `json:"grade"`
}

// NewService creates a new meteo service.
func NewService() *Service {
	return &Service{
		client:     http.DefaultClient,
		serverAddr: "http://localhost:8080",
	}
}

// getCityMeteo sends a request to retrieve meteo info for the given city.
func (s *Service) getCityMeteo(city string) (*Meteo, error) {
	resp, err := s.client.Get(fmt.Sprintf("%s/meteo/%s", s.serverAddr, city))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to perform meteo request")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read meteo response for city: %s", city)
	}

	var meteo *Meteo
	if err := json.Unmarshal(b, &meteo); err != nil {
		return nil, err
	}

	return meteo, nil
}

// GetBestCityMeteo returns the city with the best meteo.
func (s *Service) GetBestCityMeteo(cities []string) (string, error) {
	bestCity := ""
	bestScore := 10

	for _, city := range cities {
		meteo, err := s.getCityMeteo(city)
		if err != nil {
			return "", err
		}

		if meteo.Grade < bestScore {
			bestCity = city
			bestScore = meteo.Grade
		}
	}

	return bestCity, nil
}
