package main

import "fmt"

func send(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("sending ", i)
		c <- i

	}
	close(c)
}

func main() {
	c := make(chan int)
	go send(c)
	for i := range c {
		fmt.Println("received ", i)
	}
}
