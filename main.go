package main

import "fmt"

const usdEur = 0.85                //на 21.04.2026
const usdRub = 75.14               //на 21.04.2026
const eurRub = 1 / usdEur * usdRub //на 21.04.2026

func main() {

}

func getUserInput() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func convertCurrency(sum float64, originalCurrency string, targetCurrency string) float64 {
	return 0
}
