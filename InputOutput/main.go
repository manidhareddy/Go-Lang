package main

import (
	"fmt"
	"os"
	"io"
)

func main(){
	//take some set of files a command line arguments

	for _, fname := range os.Args[1:]{
		file, err := os.Open(fname)

		fmt.Printf("debug -> %T\n",file)

		if err != nil{
			fmt.Fprintln(os.Stderr,err)
			continue
		}
		if _, err := io.Copy(os.Stdout, file) ; err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		defer file.Close()
	}
}