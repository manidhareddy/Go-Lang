package main

import "fmt"

func do (z[5]int){
	z[0] = 23
	
}

func main(){
	c := [5] int{0,0,0,0}
	fmt.Printf("%T    -> %[1]v \n",c)
	do(c)
	

	fmt.Printf("%T    -> %[1]v\n ",c)
}