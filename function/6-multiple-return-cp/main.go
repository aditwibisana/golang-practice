package main

import "fmt"

//fungsi square akan mengembalikan nilai pangkat 2
//dari dua parameter yang diterima
//contoh
//parameter yang diterima (4,3)
//maka akan mengembalikan 16 dan 9
func main() {

	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
	fmt.Println(square(2, 5))

}

//gunakan * untuk melakukan perkalian
<<<<<<< HEAD

// TODO: answer here
func square(num1, num2 int) (int, int){
	result1 := num1*num1
	result2 := num2*num2
	return result1, result2

}

//beginanswer
// func square(angka1, angka2 int) (int, int) {
// 	result1 := angka1 * angka1
// 	result2 := angka2 * angka2
// 	return result1, result2
// }

//endanswer

=======
// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
