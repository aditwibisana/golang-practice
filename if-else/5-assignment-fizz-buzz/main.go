package main

import "fmt"

// Quiz ini dinamakan fizzbuzz
// Tuliskan program for loop dari 1 sampai dengan 100.
// lalu setiap perulangan program ini.

// Jika angka tersebut habis dibagi 3, maka tampilkan "fizz"
// Jika angka tersebut habis dibagi 5, maka tampilkan "buzz"
// Jika angka tersebut habis dibagi 3 dan 5, maka tampilkan "fizzbuzz"

// Tips gunakan % (modulo) untuk mengetahui apakah angka tersebut habis dibagi 3 atau 5 atau tidak.

func main() {
	for i := 1; i <= 100; i++ {
<<<<<<< HEAD
<<<<<<< HEAD
		// TODO: answer here
		if (i%3==0 && i%5==0){
			fmt.Println("fizzbuzz")
		} else if (i%5==0){
			fmt.Println("buzz")
		}else if (i%3==0){
			fmt.Println("fizz")
		}else{
			fmt.Println(i)
		}
=======
		//beginanswer
		fizz := "fizz"
		buzz := "buzz"

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(fizz + buzz)
		} else if i%3 == 0 {
			fmt.Println(fizz)
		} else if i%5 == 0 {
			fmt.Println(buzz)
		} else {
			fmt.Println(i)
		}
		//endanswer
>>>>>>> 0dbcdd8ebce63009fcee596516afbb40b893ca25
=======
		// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
	}
}
