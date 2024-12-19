package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	fmt.Println("hey, What your name?")

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	fmt.Println("hello Mr", name)
}
