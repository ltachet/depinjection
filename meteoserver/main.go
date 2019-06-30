package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/meteo/:city", getCityMeteo)
	r.Run(":8080")
}

func getCityMeteo(c *gin.Context) {
	city := strings.ToLower(c.Param("city"))

	type prediction struct {
		data string
		// grade gives the dangerosity level of the climate from 0 to 10.
		grade int
	}

	var p prediction

	switch city {
	case "paris":
		p.data = "Hot and warm"
		p.grade = 6
	case "auckland":
		p.data = "Rainy"
		p.grade = 2
	case "ottawa":
		p.data = "Snowing"
		p.grade = 5
	case "tokyo":
		p.data = "Windy"
		p.grade = 1
	default:
		c.JSON(404, gin.H{"error": "city not found"})
		return
	}

	c.JSON(200, gin.H{"data": p.data, "grade": p.grade})
}
