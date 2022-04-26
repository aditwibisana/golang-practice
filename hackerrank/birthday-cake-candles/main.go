package main

import (
	"fmt"
)

// Birthday Cake Candles
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func birthdayCakeCandles(candles []int32) int32 {
	// TODO: answer here
	//cari angka tertinggi
	//hitung ada berapa
	count:=0
	max := candles[0]
	//looping to find the max value
	for i:=0 ;i < len(candles); i++{
		if candles[i]>= max{
			max = candles[i]
		}
		// if candles[i] < max{
		// 	max = max
		// }
	}
	//looping to find the count of max value
	for _, value:= range candles{
		if value == max{
			count++
		}
	}
	//fmt.Printf("%d sebanyak ",max)
	return int32(count)	
}

func main() {
	var arr = []int32{-1, 9, 9, 8}

	fmt.Println(birthdayCakeCandles(arr))
}
