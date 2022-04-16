package main

import "fmt"

func factorialRecursive(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result

}
func main() {
	fmt.Println(factorialRecursive(5))
}
