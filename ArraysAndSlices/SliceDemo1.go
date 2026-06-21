package main

import "fmt"

func main() {
	names := make([]string, 2)

	names = append(names, "manidhar")

	for i := range names {
		fmt.Println(names[i])
	}

	other := make([]int, 0, 5)

	for i := range (2,6) {
		other = append(other, i)
	}

	for i := range other {
		fmt.Println(other[i])
	}

}
