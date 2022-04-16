package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Heroes []struct {
	Name          string `json:"name"`
	BirthYear     int    `json:"birth_year"`
	DeathYear     int    `json:"death_year"`
	Description   string `json:"description"`
	AscensionYear int    `json:"ascension_year"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/golang-heroes?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("failed to connect database", err)
	} else {
		log.Println("Connection Success")

	}
	db.AutoMigrate(&Heroes{})
	//handleRequests()
}

//func handleRequests() {
//	log.Println("Start to listening request at http://localhost:8080")
//
//	myRouter := mux.NewRouter().StrictSlash(true)
//
//	myRouter.HandleFunc("/", homePage)
//	myRouter.HandleFunc("/api/view-heroes", viewHeroes).Methods("GET")
//
//	log.Fatal(http.ListenAndServe(":8080", myRouter))
//}
//func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Welcome to my API")
//}
