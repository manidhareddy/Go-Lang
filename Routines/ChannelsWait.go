
package main

import (
	"fmt"
	"time"
)

func send(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("sending ", i)
		c <- i

	}
	
	time.Sleep(60 * time.Second);
	fmt.Println("waiting")
}

func main() {
	c := make(chan int,5)
	go send(c)
	for i:=0 ; i<5 ; i++ {
		fmt.Println("received ", <-c )
	}
}

