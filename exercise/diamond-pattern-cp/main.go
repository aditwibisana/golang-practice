package main

import "fmt"

// Check Point:
// Diamond Pattern
// - Input: Size: 5
// - Output:
//         *
//        ***
//       *****
//      *******
//     *********
//    ***********
//     *********
//      *******
//       *****
//        ***
//         *

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	// TODO: answer here
<<<<<<< HEAD
	//MYANSWER
	// for i := 1; i <= size+1; i++ {
	// 	for j := 0; j <= size-i; j++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	for k := 1; k <= i*2-1; k++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := size ; i > 0; i-- {
	// 	for j := 0; j <= size-i; j++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	for k := 1; k <= i*2-1; k++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println()
	// }

	//ENDSHERE


	//beginanswer
	var i, j, k, l int
	for i = size; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = size; k > i-1; k-- {
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i = 0; i <= size; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = size; k > i-1; k-- {
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	//endanswer

=======
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
}
