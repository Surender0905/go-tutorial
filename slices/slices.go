package main

import "fmt"

func main() {

	//slice is a dynamic array

	//slice is a reference to an array
	//slice is a reference type

	// var arr []int

	// arr1:=[]int{}

	//append is a function in go

	//length of arr

	// fmt.Println(len(arr), len(arr1), "length of dynamic arr")

	// arr = append(arr, 1)

	// arr = append(arr, 2)

	// arr = append(arr, 3)

	// fmt.Println(arr)

	// var nums=make([]int, 3, 5)

	// nums= append(nums, 4)

	// nums= append(nums, 5)

	// nums= append(nums, 6)

	// nums= append(nums, 7)

	// fmt.Println(nums)
	// //length of slice
	// fmt.Println(len(nums))
	// //capacity of slice
	// fmt.Println(cap(nums))

	// //copy slice
	// var nums2=make([]int, len(nums), cap(nums))

	// copy(nums2, nums)

	// fmt.Println(nums2, "copied slice")

	///slice operations
	var nums = []int{1, 2, 3}

	//append
	nums = append(nums, 4)

	//append
	nums = append(nums, 5)

	//insert
	nums = append(nums, 6)

	

	fmt.Println(nums)
	fmt.Println((nums[0:3]))

}