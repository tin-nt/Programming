package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//City struct
type City struct {
	ID         int
	Name       string
	Population int
}

func dbCon() (db *sql.DB) {
	db, err := sql.Open("mysql", "tinnt:123@tcp(localhost:3306)/DB_SQLi")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//Index list all city
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbCon()
	defer db.Close()

	sqlStatement := "SELECT * FROM cities ORDER BY id ASC"
	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	for res.Next() {
		var city City
		err := res.Scan(&city.ID, &city.Name, &city.Population)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%v\n", city)
	}

}

//Create insert city to db
func Create(w http.ResponseWriter, r *http.Request) {
	db := dbCon()
	defer db.Close()
	name := r.URL.Query()["name"][0]
	Population := r.URL.Query()["population"][0]

	// Filter
	// reString := regexp.MustCompile(`^[a-zA-Z]+$`)
	// reNumber := regexp.MustCompile(`^[0-9]+$`)
	// match1 := reString.MatchString(name)
	// match2 := reNumber.MatchString(Population)
	// sqlStatement := ""
	// if match1 && match2 {
	// 	sqlStatement = "INSERT INTO cities(name, population) values (?,?)"
	// } else {
	// 	http.Error(w, "Wrong pattern!", 400)
	// 	return
	// }

	res, err := db.Exec("INSERT INTO cities(name, population) values (?,?)", name, Population)

	if err != nil {
		// debug
		// log.Fatal(err)
		http.Error(w, "param missing or wrong type!", 400)
	} else {
		fmt.Fprint(w, "Inserted city into db")
		res.LastInsertId()
	}
}

//Find find city id
func Find(w http.ResponseWriter, r *http.Request) {
	db := dbCon()
	defer db.Close()
	var (
		id         string
		name       string
		population string
	)
	id = r.URL.Query()["id"][0]

	// sqlStatement := "SELECT * FROM cities WHERE id = ?"
	// res, err := db.Query(sqlStatement, id)
	res, err := db.Query("SELECT * FROM cities WHERE id = " + id)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	for res.Next() {
		err := res.Scan(&id, &name, &population)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprintf(w, "ID: %v\nCity: %v\nPopulation: %v\n", id, name, population)
	}

}

func main() {
	db := dbCon()
	defer db.Close()
	// call index
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/find", Find)
	fmt.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
