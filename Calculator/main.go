package main

import (
	"fmt"

	"example.com/calculator/loadData"
	"example.com/calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		loadData.LoadData(priceJob)
		priceJob.Process()
	}

	fmt.Println()
}
