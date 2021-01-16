package main

import (
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "tinnt:123@tcp(127.0.0.1:3306)/DB_SQLi?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		Population := c.Query("population")

		// Filter
		reString := regexp.MustCompile(`^[a-zA-Z]+$`)
		reNumber := regexp.MustCompile(`^[0-9]+$`)
		match1 := reString.MatchString(name)
		match2 := reNumber.MatchString(Population)
		if !match1 || !match2 {
			c.JSON(200, gin.H{
				"message": "Wrong pattern!" + name + Population,
			})
			return
		}

		// Create record
		PopulationInt, _ := strconv.ParseUint(Population, 10, 4)
		city := City{Id: 0, Name: name, Population: PopulationInt}
		result := db.Create(&city)
		if result.Error == nil {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Error",
			})
		}
	})
	r.Run(":2000")
}

type City struct {
	Id         uint64
	Name       string
	Population uint64
}
