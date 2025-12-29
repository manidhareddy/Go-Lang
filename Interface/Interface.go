package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

type R struct {
	P int
}

func (r R) M() {
	fmt.Println(r.P)
}

func main() {
	var i I = T{"hello"}
	i.M()

	var r I = R{10}
	r.M()
}
