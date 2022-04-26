package main

import (
	"fmt"
)

func main() {
	//membuat rectangle dengan anonymous struct
	//field dari struct ini sama seperti rectangle sebelumnya
	// TODO: answer here
	newRectangle := struct{
		Width, Length int
	}{
		Width: 15,
		Length: 2,

	}

	fmt.Printf("rectangle dengan lebar %d dan panjang %d", newRectangle.Width, newRectangle.Length)
}
