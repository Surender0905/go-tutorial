package main

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sun int = 0

	// for i := 0; i < len(num); i++ {

	// 	sun += num[i]

	// 	println(num[i])
	// }

	///now with help of range
	for i, num := range num {

		sun = sun + num

		println(num, i)
	}
	// println(sun)

	///now with maps
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}

	for key, value := range m {

		println(key, value)
	}

	///now with strings
	for i, c := range "g😊lang" {

		println(i, c, string(c))
	}
}