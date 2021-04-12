package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "tinnt"
	dbPass := "123"
	dbName := "DB_SQLi"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Index(c *gin.Context) {
	db := DBConn()
	rows, err := db.Query("SELECT * FROM cities")
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Story not found",
		})
	}
	// log.Print(rows)
	post := Post{}
	for rows.Next() {
		var id, population int
		var name string
		err = rows.Scan(&id, &name, &population)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Name = name
		post.Population = population
		c.JSON(http.StatusOK, post)

	}
	defer db.Close()
}

func Find(c *gin.Context) {
	// var city City

	// reId := regexp.MustCompile(`^[0-9]+$`)

	// if err := DB.Where("id = " + c.Query("id")).First(&city).Error; err != nil && reId.MatchString(c.Query("id")) {
	// if err := DB.Where("id = ?", c.Query("id")).First(&city).Error; err != nil {

	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": city})
	db := DBConn()
	query := "SELECT * FROM cities WHERE id = :Id"
	var Id string
	Id = c.Query("id")
	rows, err := db.Query(query, sql.Named("Id", Id))
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Story not found",
		})
	}
	// log.Print(rows)
	post := Post{}
	for rows.Next() {
		var id, population int
		var name string
		err = rows.Scan(&id, &name, &population)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Name = name
		post.Population = population
	}
	c.JSON(http.StatusOK, post)

	defer db.Close()

}

func main() {

	// dsn := "tinnt:123@tcp(127.0.0.1:3306)/DB_SQLi?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	r := gin.Default()
	r.GET("/", Index)
	r.GET("/id", Find)
	r.Run(":8081")
}
