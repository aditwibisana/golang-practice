package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.

// Print alamat memori dari masing-masing variabel dibawah ini
func main() {
	name := "Roger"
	age := 20
	isMarried := true

	// TODO: answer here
	fmt.Printf("alamat memori dari variabel name %v\n",&name)
	fmt.Printf("alamat memori dari variabel age %v\n",&age)
	fmt.Printf("alamat memori dari variabel isMarried %v\n",&isMarried)
}
