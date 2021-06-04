package main

import "fmt"

func main() {
	mapa := map[int]string{}
	maps := make(map[int]string)

	mapa[0] = "Apple"
	maps[0] = "Google"

	fmt.Println("mapa: ", mapa[0])
	fmt.Println("maps: ", maps[0])

	v, ok := maps[0]
	if ok {
		fmt.Println("Ran")
		fmt.Println(v)
	}
}
