package main

import "fmt"

func main(){
	mp := make(map[string]int)

	mp["manidhar"] = 24
	mp["rajesh"] = 23
	mp["gokul"] = 29
	mp["praveen"] = 34

	value,_ := mp["raje"]
	fmt.Println(value)
	fmt.Println(_)
}
