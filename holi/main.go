package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ltachet/depinjection/holi/api"
	"github.com/ltachet/depinjection/holi/meteo"
)

func main() {
	meteoService := meteo.NewService()
	holiHandler := api.NewHolidayHandler(meteoService)

	r := gin.Default()
	r.GET("/holiday/location", holiHandler.GetBestLocation)

	r.Run(":8081")
}
