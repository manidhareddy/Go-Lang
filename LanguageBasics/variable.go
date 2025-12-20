package main

import "fmt"

var s string = "flair"

func main() {
	var a int
	var b string
	var c float64
	var d bool

	s:="smart"

	fmt.Println(s)

	for i:=1 ; i < 3 ; i++{
		s := "oversmart"
		fmt.Println(s)
	}
	s:="overoversmart"
	fmt.Println(s)

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n\n", d, d)
	//Multiple Assignment

	j, k, l := "shark", 2.05, 15
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

}
