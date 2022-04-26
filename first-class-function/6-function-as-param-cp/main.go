package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
<<<<<<< HEAD

	// TODO: answer here
	printString := func (f func() string)  {
		result := f() //ini akan menyimpan hasil dari func yang dijadikan sbg param
		fmt.Println(result)
	}

	goodNight:= func() string{
		return "selamat malam"
	}

	//beginanswer
	// printString := func(f func() string) {
	// 	result := f()
	// 	fmt.Println(result)
	// }

	// goodNight := func() string {
	// 	return "good night friends"
	// }
	//endanswer
=======
	// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c


	printString(goodNight)

}
