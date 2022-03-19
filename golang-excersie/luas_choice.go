package main

import "fmt"

func main() {
	var (
		pil int
	)
	fmt.Println("Gudang Rumus Luas")
	fmt.Println("1. Persegi ")
	fmt.Println("2. Persegi Panjang ")
	fmt.Println("3. Segitiga ")
	fmt.Println("4. Lingkaran ")
	fmt.Printf("Masukan Pilihan Luas yang akan dihitung(1-4) : ")
	fmt.Scan(&pil)

	switch pil {
	case 1:
		var sisi int
		var luas_persegi int
		fmt.Printf("Masukan Sisi Persegi : ")
		fmt.Scan(&sisi)

		luas_persegi = sisi * sisi
		fmt.Println("Luas Persegi Adalah: \n", luas_persegi)
	case 2:
		var (
			panjang int
			lebar   int
			luas    int
		)
		fmt.Printf("Masukan Panjang : ")
		fmt.Scan(&panjang)
		fmt.Printf("Masukan Lebar : ")
		fmt.Scan(&lebar)

		luas = panjang * lebar
		fmt.Println("Luas Persegi Adalah: \n", luas)
	case 3:
		const bil = 0.5
		var (
			alas   float32
			tinggi float32
			luas   float32
		)
		fmt.Printf("Masukan Alas : ")
		fmt.Scan(&alas)
		fmt.Printf("Masukan Tinggi : ")
		fmt.Scan(&tinggi)

		luas = bil * alas * tinggi
		fmt.Printf("Luas Segitiga Adalah: %.6f\n", luas)
	case 4:
		const phi = 3.14
		var (
			jari float32
			luas float32
		)
		fmt.Printf("Masukan Jari : ")
		fmt.Scan(&jari)

		luas = phi * jari * jari
		fmt.Printf("Luas Lingkaran Adalah: %.6f\n", luas)
	}
}
