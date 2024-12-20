package main

import "fmt"

//variadic function

func sum (nums ...int) int{

	total := 0

	for _, num := range nums{
		total += num
	}

	return total


}


///variadic function with any type of data and return any type data

// func sum2 (nums ...interface{}) int{

// 	total := 0	

// 	for _, num := range nums{
// 		total += num
// 	}

// 	return total	
// }




func main (){
	//make nums slice
	

	nums := []int{1,2,3,4,5,6,7,8,9,10}
    

	//how we print like 1 2 3 4 5 as nums

	// for _, num := range nums{
	// 	fmt.Println(num)
	// }

	


	fmt.Println(nums)

	println(sum(nums...))

}