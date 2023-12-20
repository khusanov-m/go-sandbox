package main

import (
	"errors"
	"fmt"
	"os"
)

func writeInvestmentToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("investment.txt", []byte(results), 0644)
}

func main() {
	revenue, err1 := getUserValue("Revenue: ")
	// if err != nil {
	// 	panic(err)
	// }
	
	expenses, err2 :=	getUserValue("Expenses: ")
	// if err != nil {
	// 	panic(err)
	// }

	tax, err3 :=	getUserValue("Tax rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		// panic(err)
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, tax)
	writeInvestmentToFile(ebt, profit, ratio)

	fmt.Printf("Earnings before tax: %.1f\n", ebt)
	fmt.Printf("Earnings after tax: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserValue(message string) (float64, error) {
	var value float64
	fmt.Print(message)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return value, nil
}

func calculateProfit(revenue, expenses, tax float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax / 100)
	ratio = ebt / profit
	return
}