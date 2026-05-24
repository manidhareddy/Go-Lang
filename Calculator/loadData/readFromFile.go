package loadData

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/calculator/prices"
)

func LoadData(job *prices.TaxIncludedPriceJob) {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An error occurred!")
		fmt.Println(err)
		return
	}
	defer file.Close()
	//file implements Reader interface
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("while reading data", err)
			continue
		}
		job.InputPrices = append(job.InputPrices, value)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("An error occurred! while reading data")
		fmt.Println(err)
		return
	}
}
