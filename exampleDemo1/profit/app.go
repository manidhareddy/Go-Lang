package profit

import (
	"fmt"
	//"math"
)

func BaicsCalculation(){
        //Take input

	revenue := printAndScan("Enter revenue Amount : ");

	expenses := printAndScan("Enter expenses : ")

	taxRate := printAndScan("Enter tax rate : ")

        EBT := calculateEBT(revenue ,expenses)

	afterTax := EBT * (1 - ratio(taxRate,100))

	ratio := ratio(EBT,afterTax)

        fmt.Printf("EBT : %v \n",EBT);

        fmt.Printf("After Tax : %v \n",afterTax)

	fmt.Printf("Ratio : %.2f \n",ratio);



}
func printAndScan(str string) float64{
	var takeInput float64
	fmt.Print(str)
	fmt.Scan(&takeInput)
	return takeInput
}

func calculateEBT(revenue,expenses float64) float64{
	return revenue - expenses
}

func ratio(value1 float64,value2 float64) float64{

	return value1/value2;
}


