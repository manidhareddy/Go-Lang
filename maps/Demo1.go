package main

import "fmt"

func main() {
	var heap map[string]int

	heap = make(map[string]int)

	heap["macair"] = 50000
	heap["samwatch"] = 3000
	heap["car"] = 800000

	fmt.Println(heap)
}
