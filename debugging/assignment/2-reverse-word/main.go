package main

import "fmt"

func main() {
	/*
		Reverse Word:
		Example: halo -> olah
	*/
	word := "halo"
	res := ReverseWord(word)
	fmt.Println(res)

	word = "katak"
	resCorrect := ReverseWordCorrect(word)
	fmt.Println(resCorrect)
	// Try correct answer:
	// resCorrect := ReverseWordCorrect(arr)
	// fmt.Println(resCorrect)
}

func ReverseWord(word string) string {
	n := len(word)
	temp := []rune(word)

	for i := 0; i <= n; i++ {
		left := i
		right := n - i - 1
		temp[left], temp[right] = temp[right], temp[left]
	}

	return string(temp)
}

func ReverseWordCorrect(word string) string {
	n := len(word)
	temp := []rune(word)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]

	}
	return string(temp)
	// TODO: replace this
}
