package main

import (
	"fmt"
)

// Mini-Max Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/mini-max-sum/problem
// 1. buat variabel min, max, kita anggap arr[0] adl min max
// 2. buat variabel utk hasil penjumlahan
// 3. looping arr: cek value dari arr < min -> ?; value dari arr > max -> ?; hitung sum += num
// 4. fmt.Printf: sum-max, sum-min
// keyword: max sum: buang nilai terkecil, min sum: buang nilai terbesar

func miniMaxSum(arr []int32) { // dewi
	// TODO: answer here
  // anggap dulu nilai terkecil dan terbesar itu adl index pertama arr:
  min := arr[0]
  max := arr[0]
  // min, max := arr[0], arr[0]
  var sum int32
  
  for _, value := range arr{
    if value < min { // cek nilai terkecil
      min = value // nilai terkecil
    }
    if value > max {
      max = value
    }
    // sum: jumlah semua value
    sum += value

  }
  // sum terkecil: buang max
  // sum terbesar: buang min
  fmt.Println(sum-max, sum-min)
}

func main() {
	var arr = []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}