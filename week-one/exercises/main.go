package main

import "fmt"

/*
1. Find average of all the elements in a list
*/
func FindAverage(data []int) float64 {
	size := len(data)
	sum := 0
	for _, v := range data {
		sum += v
	}
	return float64(sum) / float64(size)
}

func main() {
	data := []int{2, 3, 5, 3, 6, 9, 2}
	a := FindAverage(data)
	fmt.Println("The average of the data list is ", a)

}
