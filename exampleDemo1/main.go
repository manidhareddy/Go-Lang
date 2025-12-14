package main

import (
	"exampleDemo1.com/app/profit"
	"fmt"
)

//something string = "value"

func main(){
	var something string = "value"
	fmt.Println("using some value " ,something);
	fmt.Println("profit Calculator")	
	profit.BaicsCalculation()
}
