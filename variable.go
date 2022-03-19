package main

import "fmt"

type User struct {
	name string
	age  int
	role string
}

func main() {
	// Search Role
	users := []User{
		{
			name: "Aditira",
			age:  20,
			role: "Programmer",
		},
		{
			name: "Dimas",
			age:  18,
			role: "Designer",
		},
		{
			name: "Yuni",
			age:  21,
			role: "DevOps",
		},
	}
	var searchResult []User
	var role string
	fmt.Printf("Masukan Role yang ingin dicari = ")
	fmt.Scan(&role)
	for _, user := range users {
		if user.role == role {
			searchResult = append(searchResult, user)
		}
	}
	if len(searchResult) != 0 {
		fmt.Printf("%v Found: ", role)
		for _, u := range searchResult {
			fmt.Println("Name: ", u.name, "Age: ", u.age, "Role: ", u.role)
		}
	} else {
		fmt.Printf("Role: %v Not Found", role)
	}
}

// CP Repetition with Condition
//var size int
//fmt.Printf("Masukan Jumlah Antrian = ")
//fmt.Scan(&size)
//for i := 1; i <= size; i++ {
//	if i < 6 {
//		fmt.Printf("Antrian %d Membeli beras dengan kualitas super [SUPER]\n", i)
//	} else if i > 5 && i < 11 {
//		fmt.Printf("Antrian %d Membeli beras dengan kualitas medium [MEDIUM]\n", i)
//	} else {
//		fmt.Printf("Antrian %d Membeli beras dengan kualitas low [LOW]\n", i)
//	}
//
//}

// CP Star Pattern Kanan
//var baris int
//
//fmt.Print("Masukan Jumlah Baris = ")
//fmt.Scanln(&baris)
//
//fmt.Println("\nStar Pattern")
//for i := 1; i <= baris; i++ {
//	for j := 0; j < baris-i; j++ {
//		fmt.Print(" ")
//	}
//	for j := 0; j < i; j++ {
//		fmt.Print("*")
//	}
//	fmt.Println()
//}

//var (
//	pil int
//)
//fmt.Println("Gudang Rumus Luas")
//fmt.Println("1. Persegi ")
//fmt.Println("2. Persegi Panjang ")
//fmt.Println("3. Segitiga ")
//fmt.Println("4. Lingkaran ")
//fmt.Printf("Masukan Pilihan Luas yang akan dihitung(1-4) : ")
//fmt.Scan(&pil)
//
//switch pil {
//case 1:
//	var sisi int
//	var luas_persegi int
//	fmt.Printf("Masukan Sisi Persegi : ")
//	fmt.Scan(&sisi)
//
//	luas_persegi = sisi * sisi
//	fmt.Println("Luas Persegi Adalah: \n", luas_persegi)
//case 2:
//	var (
//		panjang int
//		lebar   int
//		luas    int
//	)
//	fmt.Printf("Masukan Panjang : ")
//	fmt.Scan(&panjang)
//	fmt.Printf("Masukan Lebar : ")
//	fmt.Scan(&lebar)
//
//	luas = panjang * lebar
//	fmt.Println("Luas Persegi Adalah: \n", luas)
//case 3:
//	const bil = 0.5
//	var (
//		alas   float32
//		tinggi float32
//		luas   float32
//	)
//	fmt.Printf("Masukan Alas : ")
//	fmt.Scan(&alas)
//	fmt.Printf("Masukan Tinggi : ")
//	fmt.Scan(&tinggi)
//
//	luas = bil * alas * tinggi
//	fmt.Printf("Luas Segitiga Adalah: %.6f\n", luas)
//case 4:
//	const phi = 3.14
//	var (
//		jari float32
//		luas float32
//	)
//	fmt.Printf("Masukan Jari : ")
//	fmt.Scan(&jari)
//
//	luas = phi * jari * jari
//	fmt.Printf("Luas Lingkaran Adalah: %.6f\n", luas)
//}

//var(
//	bil1 int
//	bil2 int
//	pil int
//	hasil int
//)
//fmt.Printf("Masukan Bilangan 1: ")
//fmt.Scan(&bil1)
//fmt.Printf("Masukan Bilangan 2: ")
//fmt.Scan(&bil2)
//fmt.Println("1. Tambah (+) ")
//fmt.Println("2. Kurang (-) ")
//fmt.Println("3. Kali (*) ")
//fmt.Println("4. Bagi (/) ")
//fmt.Printf("Masukan Operator(1-4) : ")
//fmt.Scan(&pil)

// Volume Tabung
//const phi = 3.14
//var (
//	jari   float32
//	tinggi float32
//	hasil  float32
//)
//fmt.Printf("Masukan Jari-Jari: ")
//fmt.Scan(&jari)
//fmt.Printf("Masukan Tinggi: ")
//fmt.Scan(&tinggi)
//hasil = phi * jari * jari * tinggi
//fmt.Printf("Total Volume: %.6f\n", hasil)

// Luas Persegi
//var (
//	sisi float32
//	luas float32
//)
//fmt.Printf("Masukan Sisi: ")
//fmt.Scan(&sisi)
//
//luas = sisi * sisi
//fmt.Printf("Total luas: %.2f\n", luas)
