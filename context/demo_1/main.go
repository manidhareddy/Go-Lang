package main

import (
	"context"
	"fmt"
	"time"
)

// controlling timouts
// cancelling go routines
// ans passing metadata across your Go application
func main() {
	exampleTimeout()
}

func exampleTimeout() {
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 4*time.Second)

	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWithTimeout.Done():
		fmt.Println("oh no my timeout expired!", ctxWithTimeout.Err())
	}
}
