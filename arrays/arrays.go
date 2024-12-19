package main

import "fmt"

///numbered sequence of specific length

func main(){
	var nums [4]int

	nums[0] = 1
	


	println(nums[0])

	fmt.Println(nums)

    //array length
	println(len(nums))



	var val [4]string

	val[0] = "Roshan"
	

	fmt.Println(val)

	number:=[3]int{1,2,3}
	fmt.Println(number)


	///2d array


	var twoD [2][3]int
	twoD[0][0] = 1
	twoD[0][1] = 2
	twoD[0][2] = 3
	twoD[1][0] = 4
	twoD[1][1] = 5
	twoD[1][2] = 6


	fmt.Println(twoD)

	//other way of creating two d arrays

	twoD1 := [2][3]int{{1,2,3},{4,5,6}}

	fmt.Println(twoD1)

	//dynamic array creation

	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	fmt.Println(arr)
}