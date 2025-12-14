package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main(){

	for _, fname := range os.Args[1:] {

		//open file
		file, err := os.Open(fname)
		if  err != nil{
			fmt.Fprintln(os.Stderr,err)
			continue
		}

		var lc, wc, cc int

		scan := bufio.NewScanner(file)

		for scan.Scan(){
			lc ++
			s := scan.Text()
			cc+= len(s)

			wc += len(strings.Fields(s)) // seperated by space

		}

		fmt.Printf("%7d %7d %7d %s \n",lc,wc,cc,fname)
		
		file.Close()
	}
}