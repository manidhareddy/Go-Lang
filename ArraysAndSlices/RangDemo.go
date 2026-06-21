package main

import "fmt"

func main() {
	grades := [2]int{34, 45}

	for i, _ := range grades {
		fmt.Println(grades[i])
	}

}
