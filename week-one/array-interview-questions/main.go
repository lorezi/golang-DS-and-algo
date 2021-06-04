package main

import "fmt"

// Sum List --> Write a method that will return the sum of all the elements of the integer list, given list as an input argument.
func sumList(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func main() {
	data := []int{1, 4, 5, 3, 6, 9}
	sum := sumList(data)
	fmt.Println("The sum of the list is: ", sum)
}
