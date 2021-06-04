package main

import (
	"fmt"
)

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

// Binary Search function in a sorted list
func BinarySearch(list []int, val int) bool {
	size := len(list)
	var mid int
	low := 0
	high := size - 1

	for low <= high {
		// find the mid index in the list
		mid = low + (high-low)/2
		if list[mid] == val {
			return true
		} else {
			if list[mid] < val { // search right because the value is greater than the middle value
				low = mid + 1
			} else { // search left because the value is smaller than the middle value
				high = mid - 1
			}
		}
	}
	return false
}

/*

RotateArray function rotates a list in k number of times.

Analysis:
- Rotating list is done in two parts trick. In the first part, we reverse elements of list first half and then second half.
- Then we reverse the whole list, in order to complete the whole rotation
*/
func RotateArray(data []int, k int) {
	size := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, size-1)
	ReverseArray(data, 0, size-1)
}

// ReverseArray function takes list of data, start and end
func ReverseArray(data []int, start int, end int) {
	i := start
	j := end

	for i < j {
		// swapping ===> [i][j] = [j][i]
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

// Find the largest sum contiguous subarray
func MaxSubArraySum(data []int) int {

	maxSoFar := 0
	maxEndingHere := 0

	for i := range data {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}

	return maxSoFar
}

func main() {
	data := []int{1, 4, 5, 3, 6, 9}
	second_data := []int{1, 4, -5, 3, -9, 10, 11, -12, 6, 9}

	sum := sumList(data)
	fmt.Println("The sum of the list is: ", sum)

	s := sequentialSearch(data, 9)
	fmt.Println(s)

	b := BinarySearch(data, 9)
	fmt.Println(b)

	fmt.Println("Max sub array sum:", MaxSubArraySum(second_data))
}
