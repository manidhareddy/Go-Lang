package main

import "fmt"

func main() {
	var a int = 34
	if a <= 6 {
		fmt.Println("Below or equal to 6")
	} else if a <= 30 {
		fmt.Println("Below or equal to 30")
	} else {
		fmt.Println("Greater than 30")
	}
}
