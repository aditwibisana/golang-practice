package main

import (
	"fmt"
)

// A Very Big Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/a-very-big-sum/problem

func aVeryBigSum(ar []int64) int64 {
	// TODO: answer here
	var sum int64
	sum = 0
	for index := range ar {
		sum += ar[index]

	}
	return sum
}

func main() {

	var ar = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	fmt.Println(aVeryBigSum(ar))
}
