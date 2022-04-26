package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang

type Rectangle struct {
	// TODO: answer here
	Width,Length int
}

// TODO: answer here
func (r Rectangle)GetArea(){
	luas := r.Width* r.Length
	fmt.Printf("luas persegi panjang adalah %d\n", luas)
}

func(r Rectangle) GetPerimeter(){
	keliling := (2*r.Length) + (2*r.Width)
	fmt.Printf("keliling persegi panjang adalah %d", keliling)
}
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
