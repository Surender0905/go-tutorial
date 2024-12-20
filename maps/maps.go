package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	m["language"] = "Go"
	m["type"] = "backend"

	// delete(m, "type")
	// clear(m)

	// fmt.Println(m["language"], m["type"], len(m), m["test"])

	fmt.Println(m)
	testMap:=map[string]int{"one":1, "two":2, "three":3, "four":4, "five":5}
	///loop over map

	for key, value := range testMap {
		

		if(value==2){

			continue

		}

		if(value==4){

			break
		}


		fmt.Println(key, value)

	}

	///check element exists in map

	if v, ok := testMap["one"]; ok {
		fmt.Println("one exists", v)
	}else{

		fmt.Println("one does not exist", v)
	}

	///compare map
	m1:=map[string]int{"one":1, "two":2, "three":3, "four":4, "five":5}
	m2:=map[string]int{"one":1, "two":2, "three":3, "four":4, "five":7}

	fmt.Println(maps.Equal(m1, m2))
	

}