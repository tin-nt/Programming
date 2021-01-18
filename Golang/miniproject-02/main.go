package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type City struct {
	ID         int
	Name       string
	Population int
}

var DB *gorm.DB

func Index(c *gin.Context) {
	var cities []City
	DB.Find(&cities)
	c.JSON(http.StatusOK, gin.H{"data": cities})
}

func Find(c *gin.Context) {
	var city City

	// reId := regexp.MustCompile(`^[0-9]+$`)

	// if err := DB.Where("id = " + c.Query("id")).First(&city).Error; err != nil && reId.MatchString(c.Query("id")) {
	if err := DB.Where("id = ?", c.Query("id")).First(&city).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": city})

}
func main() {

	dsn := "tinnt:123@tcp(127.0.0.1:3306)/DB_SQLi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}
	db.AutoMigrate(&City{})
	DB = db
	r := gin.Default()

	r.GET("/", Index)
	r.GET("/id", Find)
	r.Run()
}
