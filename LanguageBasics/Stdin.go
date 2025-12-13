package main
import (
	"fmt"
	"os"
)
/*
 go run Stdin.go  < nums.txt
 cat nums.txt | go run Stdin.go
 */
func main() {
	var sum float64
	var n int

	for {
		var val float64
		_,err:=fmt.Fscanln(os.Stdin,&val)//Fscanln which will scan a line of text ending with new line

		if err != nil {
			break // exit for loop
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Println(os.Stderr,"empty") // print error
		os.Exit(-1) // exit program with return status -1
	}

	fmt.Println("Average numbers : ",sum/float64(n))
}
