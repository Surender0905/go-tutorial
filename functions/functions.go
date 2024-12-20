package main

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages()(string, string,string){
// 	return "golang", "java", "python"
// }

///function take as argument
func processIt(fn func(a int, b int) int) int {
	return fn(1, 2)
}

///function as return
func returnIt() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}
func main() {

	// println(add(1, 2))
	// lang1, lang2, lang3 := getLanguages()

	// languages := lang1 + " " + lang2 + " " + lang3	

	// println(languages)
	fn:=func (a int, b int) int{
		return a+b
	}
	println(processIt(fn))

	fn1:=returnIt()
	println(fn1(1,2))

	
}