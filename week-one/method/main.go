package main

import (
	"fmt"
)

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type MyInt int

func (data MyInt) increment1() {
	data = data + 1
}

func (data *MyInt) increment2() {
	*data = *data + 1
}

func main() {
	r := Rect{width: 10, height: 10}
	fmt.Println("Area: ", r.Area())
	fmt.Println("Perimeter: ", r.Perimeter())

	ptr := &Rect{width: 10, height: 5}
	fmt.Println("Area: ", ptr.Area())
	fmt.Println("Perimeter: ", ptr.Perimeter())

	var data MyInt = 1
	fmt.Println("value before increment1() call: ", data)
	data.increment1()
	fmt.Println("value after increment1call() call : ", data)
	data.increment2()
	fmt.Println("value after increment2() call : ", data)
}
