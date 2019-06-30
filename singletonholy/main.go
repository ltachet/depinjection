package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/holiday/location", getBestHolidayLocation)
	r.Run(":8081")
}

func getBestHolidayLocation(c *gin.Context) {
	bestPlace, err := getBestHolidayPlace()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": bestPlace})
}
