package main

import "fmt"

func main() {
	value := 20

	var ptr *int = &value

	fmt.Printf("Value is %d , and address is %p\n", value, ptr)

	*ptr = 40

	fmt.Println("value ", value)
}
