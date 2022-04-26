package main

import "fmt"

// Check Point:
// Start Pattern
// - Input: Size
// - Output: Start Pattern

// Contoh:
// - Input: Size: 10
// - Output:
//           *
//          **
//         ***
//        ****
//       *****
//      ******
//     *******
//    ********
//   *********
//  **********

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

<<<<<<< HEAD

	// TODO: answer here
	//MYANSWER
	// for i := 1; i <= size; i++{
	// 	for j:= i; j< size; j++{
	// 		fmt.Print(" ")
	// 	} 
	// 	for k:=1; k<=i; k++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	//ENDSHERE

	//beginanswer
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
	//endanswer

=======
	// TODO: answer here
>>>>>>> 2ae988a7ab8fa2ea4f367e7c567e39c5bce0995c
}

