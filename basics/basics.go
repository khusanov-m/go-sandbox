package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// investmentAmount := 1000
	// expectedReturnRate := 5.5
	// years := 10

	// futureValue := float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years)) 

	// fmt.Println("Investment Amount: ", futureValue)

	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)


	// fmt.Println("Future Amount: ", futureValue) 
	// fmt.Println("Future Real Amount: ", futureRealValue)

	// fmt.Printf("Future Amount: %.2f\n", futureValue)
	// fmt.Printf("Future Real Amount: %.2f\n", futureRealValue)

	formattedFV := fmt.Sprintf("Future Amount: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Amount: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func calculateFutureValues(investmentAmount , expectedReturnRate , years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
	// return // just return will be enough if the return values are named
}