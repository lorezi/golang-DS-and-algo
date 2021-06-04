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

// Sequential Search ---> Write a method, which will search a list for some given value.
func sequentialSearch(list []int, val int) string {
	for _, v := range list {
		if v == val {
			return fmt.Sprintf("Found the value in the list %d", val)
		}
	}
	return fmt.Sprintf("Value not found %d", val)
}

func main() {
	data := []int{1, 4, 5, 3, 6, 9}
	sum := sumList(data)
	fmt.Println("The sum of the list is: ", sum)

	s := sequentialSearch(data, 9)
	fmt.Println(s)
}
