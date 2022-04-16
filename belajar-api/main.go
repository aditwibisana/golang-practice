package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Heroes struct {
	Name          string `json:"name"`
	BirthYear     string `json:"birth_year"`
	DeathYear     string `json:"death_year"`
	Description   string `json:"description"`
	AscensionYear string `json:"ascension_year"`
}
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/golang-heroes?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("failed to connect database", err)
	} else {
		log.Println("Connection Success")

	}
	db.AutoMigrate(&Heroes{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start to listening request at http://localhost:8080")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/view-heroes", viewHeroes).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my API")
}
func viewHeroes(w http.ResponseWriter, r *http.Request) {
	linkUrl := ("https://indonesia-public-static-api.vercel.app/api/heroes")
	resp, err := http.Get(linkUrl)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	var heroes []Heroes
	if err := json.NewDecoder(resp.Body).Decode(&heroes); err != nil {
		fmt.Println(err.Error())
		return
	}
	res := Result{
		Code:    http.StatusOK,
		Data:    heroes,
		Message: "Success",
	}
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
