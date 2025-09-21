package main

import "fmt"

func main() {
	var a int = 10
	switch a {
	case 20:
		fmt.Println("20")
	case 10:
		fmt.Println("10")
	case 60:
		fmt.Println("60")
	default:
		fmt.Println("No match")
	}
}
