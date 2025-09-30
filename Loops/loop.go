package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Println("tick")
		time.Sleep(1 * time.Second)
		break
	}
}
