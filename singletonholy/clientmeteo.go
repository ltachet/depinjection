package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/pkg/errors"
)

var (
	httpClient *http.Client
	once       sync.Once
	cities     = []string{"paris", "auckland", "tokyo", "ottawa"}
)

// GetHTTPClient is a singleton returning an http client instance.
func GetHTTPClient() *http.Client {
	once.Do(func() {
		httpClient = http.DefaultClient
	})

	return httpClient
}

type Meteo struct {
	Data  string `json:"data"`
	Grade int    `json:"grade"`
}

func getBestHolidayPlace() (string, error) {
	c := GetHTTPClient()
	bestCity := ""
	bestScore := 10

	for _, city := range cities {
		resp, err := c.Get(fmt.Sprintf("http://localhost:8080/meteo/%s", city))
		if err != nil {
			return "", errors.Wrapf(err, "failed to get meteo for city: %s", city)
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.Wrapf(err, "failed to read meteo response for city: %s", city)
		}

		var meteoResp Meteo
		if err := json.Unmarshal(b, &meteoResp); err != nil {
			return "", err
		}

		//

		if meteoResp.Grade < bestScore {
			bestCity = city
			bestScore = meteoResp.Grade
		}
	}

	return bestCity, nil
}
