package investments

import (
	"fmt"
	"math"
)

func InvestmentValues(){
	const  infulationRate float64 = 2.5
	var investmentAmount float64 //take input from user
	var years float64
	var returnRate float64

	//Take input

	fmt.Print("Enter investment Amount : ")
	fmt.Scan(&investmentAmount) //pass address

	fmt.Print("Enter years : " )
	fmt.Scan(&years);

	fmt.Print("Enter return rate : ")
	fmt.Scan(&returnRate);

	futureValue := investmentAmount * math.Pow(1 + returnRate/100,years)

	futureRealValue := futureValue / math.Pow(1 + infulationRate/100,years)

	fmt.Printf("Future Total returns : %v \n",futureValue);

	fmt.Printf("Infulation Future Total returns : %v \n",futureRealValue);
	
}
