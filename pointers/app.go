package main

import "fmt"

func main(){
	value := 20

	ptr := &value;

	*ptr = 40

	fmt.Println("value " ,value)
}
