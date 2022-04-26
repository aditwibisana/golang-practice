package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
<<<<<<< HEAD

	// TODO: answer here
	getAreaFunc := func () func(int, int) int {
		return func(num1, num2 int) int {
			return num1 *num2
		}
	}

	//beginanswer
	// getAreaFunc := func() func(int, int) int {
	// 	return func(x, y int) int {
	// 		return x * y
	// 	}
	// }
	//endanswer

=======
	// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
	areaF := getAreaFunc()
	res := areaF(3, 4) // 12
	fmt.Println(res)

	fmt.Println(areaF(4,5))
}
