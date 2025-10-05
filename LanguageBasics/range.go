package main

import "fmt"

func main() {
	var names []string = []string{"martin", "steve", "anand", "manidhar"}
	//Go is very strict about unused variables. If you declare a variable but never use it, the compiler will throw an error.

	//   _ is a blank identifier → tells Go you don’t care about this value.
	for _, n := range names {
		fmt.Println(n)
	}
}
