package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// method
func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxAbleAmount := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxAbleAmount)
	}
	fmt.Println(result)
}

// return a pointer
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
func (job *TaxIncludedPriceJob) LoadData() {
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
