package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, cobalah untuk mempraktikkan data evaluation pada template.
// Lengkapi function ExecuteToByteBuffer sehingga objek dari struct Account dapat ter-render ke dalam template.
// Gunakan bytes.Buffer sebagai io.Writer pada template.
// Lengkapi juga textTemplate, sehingga variabel dari objek Account dapat ter-render ke dalam template.
// Contoh:
// acc := {Name: "Tony", Number: "1002321", Balance: 1000}
// Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000.

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error){
	var textTemplate string
  // tmpl, err := template.New("").Parse(textTemplate)
	// TODO: answer here
	// 1. isi dulu textTemplate
	// 2. buat template baru, parsing textTemplate, gunakan bytes.Buffer
	textTemplate = `Akun {{.Name}} dengan Nomor {{.Number}} memiliki saldo sebesar ${{.Balance}}.`
	tmpl, err := template.New("account").Parse(textTemplate)
	if err !=nil{
		return nil, err
	}
  
	var a bytes.Buffer
	err = tmpl.Execute(&a, account)
	if err !=nil{
		return nil, err
	}
  
  return a.Bytes(), nil
}

func main() {
	account := Account{
		Name:    "Marwan",
		Number:  "12345",
		Balance: 1000000,
	}

	res, _ := ExecuteToByteBuffer(account)
	fmt.Println(string(res))

}
