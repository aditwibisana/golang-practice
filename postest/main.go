package main

import "fmt"

//func addAll(nums ...int) int {
//	total := 0
//	for _, v:= range nums{
//		total += v
//	}
//	return total
//}
//func main()  {
//	points:= []int{5,10,15}
//	fmt.Println(addAll(points...))
//}
//func main() {
//	var i, j, k int
//	for i = 11; i > 0; i-- {
//		for j = 1; j <= i; j++ {
//			fmt.Printf(" ")
//		}
//		for k = 11; k > i; k-- {
//			fmt.Printf("#")
//		}
//		fmt.Printf("\n")
//	}
//}
//unc main() {
//	for bil := 1; bil < 50; bil++ {
//		i := 0
//		for bag := 1; bag < 50; bag++ {
//			if bil%bag == 0 {
//				i++
//			}
//		}
//		if (i == 2) && (bil != 1) {
//			fmt.Println(bil)
//		}
//	}
//
func main() {
	animals := [2]string{"birds", "fish"}
	fmt.Println(len(animals))
}
