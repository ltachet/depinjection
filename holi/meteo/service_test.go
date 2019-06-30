package meteo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBestCityMeteo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	}))
	defer server.Close()

	s := Service{client: server.Client(), serverAddr: server.URL}

	s.GetBestCityMeteo( /*...*/ )
}
