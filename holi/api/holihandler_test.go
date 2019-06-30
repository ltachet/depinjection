package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type meteoMock struct {
	mock.Mock
}

func (m *meteoMock) GetBestCityMeteo(cities []string) (string, error) {
	args := m.Called(cities)
	return args.String(0), args.Error(1)
}

func TestGetBestLocation(t *testing.T) {
	tests := []struct {
		// test name
		name string
		// everything needed for mocking
		citiesMock []string
		retCity    string
		retErr     error
		// final result expected
		expectedCode int
		expectedData string
	}{
		{
			name:         "working meteo service",
			citiesMock:   []string{"paris", "auckland", "tokyo", "ottawa"},
			retCity:      "auckland",
			retErr:       nil,
			expectedCode: http.StatusOK,
			expectedData: `{"data":"auckland"}`,
		},
		{
			name:         "not working meteo service",
			citiesMock:   []string{"paris", "auckland", "tokyo", "ottawa"},
			retCity:      "",
			retErr:       errors.New("service unavailable"),
			expectedCode: http.StatusInternalServerError,
			expectedData: `{"error":"service unavailable"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			meteoSvcMock := &meteoMock{}
			meteoSvcMock.On("GetBestCityMeteo", test.citiesMock).Return(test.retCity, test.retErr).Once()

			s := NewHolidayHandler(meteoSvcMock)
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			s.GetBestLocation(ginCtx)

			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedData, w.Body.String())
			meteoSvcMock.AssertExpectations(t)
		})
	}
}
