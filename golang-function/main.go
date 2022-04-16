package main

import "fmt"

//func sayHello() {
//	fmt.Println("Bayu")
//
//}
//func main() {
//	sayHello()
//}
//func IsEvenNumber(number int) bool {
//	if number%2 == 0 {
//		return true
//	}
//	return false
//}
//func main() {
//	isEvenNumber := IsEvenNumber(11)
//	fmt.Println(isEvenNumber)
//}
func multiply(multiplier int) func(x int) int {
	return func(x int) int {
		return x * multiplier
	}
}

func main() {
	multiplyByTwo := multiply(2)
	multiplyByThree := multiply(3)

	fmt.Println(multiplyByTwo(2))   // 4
	fmt.Println(multiplyByThree(2)) // 6
}
