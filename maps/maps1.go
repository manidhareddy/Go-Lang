package main

import "fmt"

func main(){
	
	var m map[string]int //nil, no storage

	//m["jon"] = 90 // PANIC - nil map

	var p map[string]int = make(map[string]int) // non-nil but empty

	p["jon"] = 90

	a := m["jon"] // return 0

	fmt.Printf("%T -> %[1]v\n",a)

	m = p
	
	m["jon"] = 120

	fmt.Println(p)

	fmt.Println(m)
}
