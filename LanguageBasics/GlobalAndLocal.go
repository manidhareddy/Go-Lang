package main

import "fmt"

var num1 = 5 //global variable

func printNumbers() {
	fmt.Println(num1)

	num1 := 10 //local variable
	num2 := 7

	fmt.Println(num1)
	fmt.Println(num2)
}

func main() {
	fmt.Println(num1)
	num1 = 56
	printNumbers()
	fmt.Println(num1)
}
