package main

import (
		"fmt"
		"strings"
       )

func main(){
	var oldText string = "ravi is a good boy ravi likes cars ravi loves"
	old := "ravi"
	new := "rana"
	s := strings.Split(oldText,old)
	
	fmt.Println(s)
	//split space
	words := strings.Split(oldText," ")

	for _,i := range words{
		fmt.Println(i)
	}

	fmt.Println("Words : ",words)

	t := strings.Join(s,new)

	fmt.Println(t)

	//join with \t 

	val := strings.Join(words,"\t");

	fmt.Println(val)
}
