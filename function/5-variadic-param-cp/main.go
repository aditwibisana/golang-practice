package main

import "fmt"

//fungsi printWord akan melakukan print satu persatu nilai parameter yang diterimanya
//contoh nilai parameter yang diterima
//("selamat","pagi","siang",sore)
//outputnya
//selamat
//pagi
//siang
//sore
func main() {
	printWord("halo")
	printWord("halo", "selamat siang")
	printWord("halo", "andi", "dan", "bagus")
	printWord("mencoba", "variadic", "param", "pada", "go")
}

<<<<<<< HEAD

// TODO: answer here
func printWord(words ...string ){
	result := ""

	for _, i :=range words{
		result += " " + i
	}

	fmt.Println(result)
}

//beginanswer
// func printWord(words ...string) {
// 	for _, word := range words {
// 		fmt.Println(word)
// 	}
// }

//endanswer

=======
// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c