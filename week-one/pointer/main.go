package main

import "fmt"

// Pointers are nothing more than variables that store memory addresses of another variable and can be used to access the value stored at those addresses

func main() {
	data := 10
	ptr := &data
	fmt.Println("Value stored at variable var is ", data)
	// deferencing a pointer type
	fmt.Println("Value stored at variable var is ", *ptr)
	fmt.Println("The address of variable var is ", &data)
	fmt.Println("The address of variable var is ", ptr)
}
