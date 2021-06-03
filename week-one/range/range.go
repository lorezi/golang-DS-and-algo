package main

import "fmt"

func arrays_and_maps_and_strings() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i, v := range numbers {
		sum += v
		fmt.Print("[", i, ",", v, "]")
	}

	fmt.Println("\nSum is :: ", sum)
	// maps
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, "-> ", v)
	}

	str := "Hello, World!"
	for i, c := range str {
		fmt.Print("[", i, ",", string(c), "]")
	}

}

func main() {
	arrays_and_maps_and_strings()
}
