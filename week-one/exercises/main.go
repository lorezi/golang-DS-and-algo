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

/*
2. Find the sum of all the elements of a two dimensional list
*/
func Sum2DList(data [][]int) float64 {
	sum := 0
	c := 0
	for _, r := range data {
		for _, v := range r {
			sum += v
			c++
		}
	}
	return float64(sum) / float64(c)
}

func main() {
	data := []int{2, 3, 5, 3, 6, 9, 2}
	a := FindAverage(data)
	fmt.Printf("The average of the data list is %.2f\n", a)

	two_dim := [][]int{{1, 2, 34}, {1, 2, 34}}
	ans := Sum2DList(two_dim)

	fmt.Printf("The sum of the 2D array is %.2f\n", ans)

}
