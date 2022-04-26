package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Dari contoh sebelumnya
// buat JSON string array nested seperti berikut
/*
[
{
		"jenis": "Meja Makan",
		"warna": "Coklat",
		"jumlah": 20,
		"ukuran": {
				"panjang": "50 cm",
				"tinggi": "25 cm"
		}
},
{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
}
]
*/

type Ukuran struct {
	// TODO: answer here
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	// TODO: answer here
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	// assign field name `Ukuran` dengan
	// struct `Ukuran` yang sudah dibuat sebelumnya
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
	jsonData, err := json.Marshal(m.MejaMeja)
	if err != nil{
		log.Println(err)
	}
	return string(jsonData)
}

func NewMeja(m Items) Items {
	return m
}

func main(){
	mejaMeja := []Meja{
			{
				Jenis:  "Meja Makan",
				Warna:  "Coklat",
				Jumlah: 20,
				Ukuran: Ukuran{
					Panjang: "50 cm",
					Tinggi:  "25 cm",
				},
			},
			{
				Jenis:  "Meja Lipat",
				Warna:  "Hitam",
				Jumlah: 1,
				Ukuran: Ukuran{
					Panjang: "70 cm",
					Tinggi:  "30 cm",
				},
			},
	
	}

	items := Items{
		MejaMeja: mejaMeja,
	}

	fmt.Println(items.EncodeJSON())



}
