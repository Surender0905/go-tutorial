package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch statement in go

	var role = "admin"

	switch role {

	case "admin":
		fmt.Println("admin user")
	case "moderator":
		fmt.Println("moderator user")
	default:
		fmt.Println("unknown user")
	}

	///multiple conditions switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}


	///type switch

	whoAmI:= func(i interface{}){
		switch i := i.(type) {

		case int:
			fmt.Println("int", i)
		case string:
			fmt.Println("string", i)
		default:
			fmt.Println("unknown", i)
		}
	}

	whoAmI(1)
	whoAmI("roshan")
	whoAmI(true)	//unknown
	
}