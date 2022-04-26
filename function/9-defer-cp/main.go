package main

import "fmt"

func main() {
	order("teh")
	order("kopi")
}

func order(drink string) {
	// mengucapkan terima kasih di akhir dengan menggunakan defer
<<<<<<< HEAD

	// TODO: answer here
	defer fmt.Println("Terima kasih sudah order")

	// //beginanswer
	// defer fmt.Println("terima kasih, silahkan datang kembali")
	// //endanswer

=======
	// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
	fmt.Println("kami sedang ada diskon untuk pembelian kopi")
	fmt.Println("pesanan anda:", drink)
	if drink == "kopi" {
		fmt.Println("okee, karena diskon totalnya jadi 4000")
	}
	fmt.Println("mohon ditunggu")

	
}
