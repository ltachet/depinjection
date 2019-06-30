package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ltachet/depinjection/holi/meteo"
)

// HolidayHandler contains all the necessary services to retrieve the best holiday location.
type HolidayHandler struct {
	meteoSvc meteo.Interface
}

// NewHolidayHandler creates a new Holiday handler.
func NewHolidayHandler(meteoSvc meteo.Interface) *HolidayHandler {
	return &HolidayHandler{
		meteoSvc: meteoSvc,
	}
}

// GetBestLocation returns the best holiday location.
func (hh *HolidayHandler) GetBestLocation(c *gin.Context) {
	var cities = []string{"paris", "auckland", "tokyo", "ottawa"}

	bestCity, err := hh.getBestLocation(cities)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": bestCity})
}

// getBestLocation returns the best holiday location.
func (hh *HolidayHandler) getBestLocation(cities []string) (string, error) {
	//var bestLocation string
	bestMeteoLocation, err := hh.meteoSvc.GetBestCityMeteo(cities)

	// Complex algo

	return bestMeteoLocation, err
}
