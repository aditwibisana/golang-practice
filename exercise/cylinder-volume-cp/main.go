package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
<<<<<<< HEAD
	//myanswer
	// var rad, height, result float64
	// fmt.Print("Masukkan nilai jari-jari alas tabung: ")
	// fmt.Scan(&rad)
	// fmt.Print("Masukkan nilai tinggi tabung: ")
	// fmt.Scan(&height)

	// result = math.Pi * (math.Pow(rad, 2)) * height

	// fmt.Println("Jadi volume tabung tersebut adalah: ", result)
	//endshere
	

	//beginanswer
	var (
		r      float32
		height float32
		volume float32
	)
	const pi = 3.14
	fmt.Printf("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)
	fmt.Printf("Masukkan tinggi tabung : ")
	fmt.Scan(&height)
	volume = pi * r * r * height
	fmt.Printf("Jadi volumenya adalah : %f\n", volume)
	//endanswer

=======
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
}
