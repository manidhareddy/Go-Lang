package main

import "fmt"

func main() {

	for i, j := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i, j)
	}
}
