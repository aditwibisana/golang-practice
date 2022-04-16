package main

import "fmt"

func sumAll(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func main() {
	//var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum = sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum)
}
