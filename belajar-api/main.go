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
	BirthYear     string `json:"birth_year,int"`
	DeathYear     int    `json:"death_year"`
	Description   string `json:"description"`
	AscensionYear int    `json:"ascension_year"`
}
type Gunung struct {
	Nama                    string `json:"nama"`
	Bentuk                  string `json:"bentuk"`
	TinggiMeter             string `json:"tinggi_meter"`
	EstimasiLetusanTerakhir string `json:"estimasi_letusan_terakhir,int"`
	Geolokasi               string `json:"geolokasi"`
}

type Hewan struct {
	Nama    string `json:"nama"`
	Makanan string `json:"makanan"`
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
	db.AutoMigrate(&Heroes{}, &Gunung{}, &Hewan{})
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
	linkUrl := ("https://mocki.io/v1/3caa2d1a-ac48-4e5a-b8e8-db66d4e698dd")
	resp, err := http.Get(linkUrl)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	var hewan []Hewan
	if err := json.NewDecoder(resp.Body).Decode(&hewan); err != nil {
		fmt.Println(err.Error())
		return
	}
	res := Result{
		Code:    http.StatusOK,
		Data:    hewan,
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

	for _, v := range hewan {
		db.Create(&v)
	}

}
