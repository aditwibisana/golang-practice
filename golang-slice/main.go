package main

import "fmt"

type User struct {
	name string
	age  int
	role string
}

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
