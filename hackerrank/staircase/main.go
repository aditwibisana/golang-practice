package main

import (
	"fmt"
)

// Staircase
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/staircase/problem

func staircase(n int32) {
	// TODO: answer here
	var i, j, k int32
	for i = 1; i<= n; i++{
		for j = i; j < n; j++{
			fmt.Print(" ")
		}
		for k = 1; k<=i; k++{
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	staircase(10)
}
