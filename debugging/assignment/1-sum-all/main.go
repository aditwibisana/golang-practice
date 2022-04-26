package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	resCorrect := SumAllCorrect(arr)
	fmt.Println(resCorrect)

	res := SumAll(arr)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := SumAllCorrect(arr)
	// fmt.Println(resCorrect)
}

func SumAll(arr []int) int {
	res := 0
	for val := range arr {
		res += val
	}
	return res
}

func SumAllCorrect(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	resCorrect := 0
	for _, val := range arr {
		resCorrect += val
	}
	return resCorrect
	// TODO: replace this
}
