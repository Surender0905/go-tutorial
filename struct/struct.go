package main

import (
	"fmt"
	"time"
)

//order of declaration
type order struct{
	
	id int
	name string
	amount float64
	status string
	createdAt time.Time  
}

func newOrder (id int, name string, amount float64) *order{

	return &order{
		id: id,
		name: name,
		amount: amount,
		status: "pending",
		createdAt: time.Now(),
	}
} 


func(o *order) changeStatus(status string ){

	o.status = status

}


func main(){

	order1 := newOrder(1, "roshan", 1000.00)

	// order1 := order{
	// 	id: 1,
	// 	name: "roshan",
	// 	amount: 1000,
	// 	status: "pending",
	// 	// createdAt: time.Now(),
	// }

	//create a new order
	// order2 := order{
	// 	id: 2,
	// 	name: "roshan",
	// 	amount: 500.00,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// order1.createdAt = time.Now()

	// fmt.Println(order1)

	// fmt.Println("order2",order2)

	// //get value from order1
	// fmt.Println(order1.id)
	// fmt.Println(order1.name)
	// fmt.Println(order1.amount)
	// fmt.Println(order1.status)
	// fmt.Println(order1.createdAt)

	// order1.changeStatus("confirmed")

	fmt.Println(order1)

}