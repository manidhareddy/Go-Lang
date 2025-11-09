package profit

import (
	"fmt"
	//"math"
)

func BaicsCalculation(){
        var revenue float64 //take input from user
        var expenses float64
        var taxRate float64

        //Take input

        fmt.Print("Enter revenue Amount : ")
        fmt.Scan(&revenue)//pass address

        fmt.Print("Enter expenses : " )
        fmt.Scan(&expenses);

        fmt.Print("Enter tax rate : ")
        fmt.Scan(&taxRate);

        EBT := calculateEBT(revenue ,expenses)

	afterTax := EBT * (1 - ratio(taxRate,100))

	ratio := ratio(EBT,afterTax)

        fmt.Printf("EBT : %v \n",EBT);

        fmt.Printf("After Tax : %v \n",afterTax)

	fmt.Printf("Ratio : %v \n",ratio);



}

func calculateEBT(revenue,expenses float64) float64{
	return revenue - expenses
}

func ratio(value1 float64,value2 float64) float64{

	return value1/value2;
}


