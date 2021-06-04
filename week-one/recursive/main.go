package main

import "fmt"

/*
Factorial calculation N! = N*(N-1)... 2*1

Where factorial of 5 is 1 * 2 * 3 * 4 * 5 = 120

*/
func Factorial(n int) int {
	// Termination condition
	if n <= 1 {
		return 1
	}

	// Body (recursive expansion)
	return n * Factorial(n-1)

}

func main() {
	n := 1
	f := Factorial(n)
	fmt.Printf("Factorial of %v is %v", n, f)
}
