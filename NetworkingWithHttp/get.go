package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	//"io"
)

var val = []string{"https://www.google.com"}

func main() {
	for _, url := range val {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		// 
		fmt.Println(resp.Status)
		resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// // fmt.Printf("%s", b)


	}
}
