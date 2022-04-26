package main

import "fmt"

func main() {
	//fungsi counter akan menerima (x int) dan mengembalikan fungsi
	//fungsi yang dikembalikan akan melakukan decrement nilai parameter x ketika dipanggil dan
	//mengembalikan nilai parameter x

	counter := func(x int) func() int {
<<<<<<< HEAD

		// TODO: answer here
		return func() int {
			x--
			return x 
		}

		//beginanswer
		// return func() int {
		// 	x--
		// 	return x
		// }
		//endanswer

=======
		// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
	}
	decrement := counter(5)
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
