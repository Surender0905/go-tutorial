package main

import "fmt"

// func changeNum(num int) {
// 	num = 5

// 	fmt.Println("in change num", num)

// }

//by reference

func changeNum(num *int) {
	*num = 5

	fmt.Println("in change num", *num)

}

func main() {
	num:= 1

	// changeNum(num)

	fmt.Println("memory address of num", &num)

	changeNum(&num)

	fmt.Println("in main", num)

}