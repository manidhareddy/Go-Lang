package main

import "fmt"

func do (z[5]int){
	z[0] = 23
	
}

func main(){
	c := [...] int{0,0,0,0}
	fmt.Printf("%T    -> %[1]v  , %d\n",c,cap(c))
//	do(c)
	

	fmt.Printf("%T    -> %[1]v\n ",c)
}
