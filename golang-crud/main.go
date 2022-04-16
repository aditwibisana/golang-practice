package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

var db *gorm.DB
var err error

// Prodcut is representation of product table
type Product struct {
	ID    int             `json:"id"`
	Code  string          `json:"code"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price" sql:"type:decimal(16,2)"`
}
type Recipes []struct {
	Calories    string `json:"calories"`
	Carbos      string `json:"carbos"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	Fats        string `json:"fats"`
	Headline    string `json:"headline"`
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Proteins    string `json:"proteins"`
	Thumb       string `json:"thumb"`
	Time        string `json:"time"`
}

// Result is an array of Product
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/golang-api?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	} else {
		log.Println("Connection Success")
	}
	db.AutoMigrate(&Product{})
	handleRequests()
}
func handleRequests() {
	log.Println("Start to listening request at http://localhost:8080")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/products", createProducts).Methods("POST")
	myRouter.HandleFunc("/api/get-products", getProducts).Methods("GET")
	myRouter.HandleFunc("/api/view", createProductsJson).Methods("POST")
	myRouter.HandleFunc("/api/view-recipes", viewRecipesJson).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my API")
}

func createProducts(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(payloads, &product)

	db.Create(&product) // Create Product

	res := Result{
		Code:    http.StatusCreated,
		Data:    product,
		Message: "Success",
	}
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)

}

func getProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	db.Find(&products)

	res := Result{
		Code:    http.StatusOK,
		Data:    products,
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
func createProductsJson(w http.ResponseWriter, r *http.Request) {
	linkUrl := ("https://s3.us-west-2.amazonaws.com/secure.notion-static.com/93e51433-0b87-4cbc-a7bd-c6b9041a03bb/recipes.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220407%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220407T163524Z&X-Amz-Expires=86400&X-Amz-Signature=b18062fb28a79af20e168653a5821449c32a676d1f0c9cd186dc0ece736e5032&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22recipes.json%22&x-id=GetObject")
	resp, err := http.Get(linkUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var recipes Recipes
	json.Unmarshal(body, &recipes)
	db.Create(&recipes)
	res := Result{
		Code:    http.StatusCreated,
		Data:    recipes,
		Message: "Success",
	}
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)

}
func viewRecipesJson(w http.ResponseWriter, r *http.Request) {
	linkUrl := ("https://s3.us-west-2.amazonaws.com/secure.notion-static.com/93e51433-0b87-4cbc-a7bd-c6b9041a03bb/recipes.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220407%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220407T163524Z&X-Amz-Expires=86400&X-Amz-Signature=b18062fb28a79af20e168653a5821449c32a676d1f0c9cd186dc0ece736e5032&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22recipes.json%22&x-id=GetObject")
	resp, err := http.Get(linkUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var recipes Recipes
	json.Unmarshal(body, &recipes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//w.Write(body)
	fmt.Println((recipes)[0])
}
