package main

import (
	"fmt"
)

func main(){
	var firstName string = "jon"
	var lastName string = "snow"

	fullName := firstName + " " + lastName

	fmt.Println("Name : " ,fullName)

	bol := fullName[5:]

	fmt.Println("lastName : ",bol)

}
