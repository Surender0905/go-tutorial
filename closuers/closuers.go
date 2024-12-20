package main

//closure function


func closure() func()int{

	var count int=0

	return func()int{

		count++
		return count

	}

}

func main() {

	c:=closure()
	println(c())
	println(c())
	println(c())

}