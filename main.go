package main

import (
	"errors"
	"fmt"
)

const usdEur = 0.85                //на 21.04.2026
const usdRub = 75.14               //на 21.04.2026
const eurRub = 1 / usdEur * usdRub //на 21.04.2026
const eurUsd = 1.17
const rubUsd = 0.013
const rubEur = 0.011

func main() {
	for {
		fmt.Println("*Конвертер валют*")
		cur, err := getOriginalCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum, err := getSum()
		if err != nil {
			fmt.Println(err)
			continue
		}
		targ, err := getTargetCurrency(cur)
		if err != nil {
			fmt.Println(err)
			continue
		}
		total := convertCurrency(sum, cur, targ)
		fmt.Println(total)
		restart := wantToRestart()
		if restart != true {
			break
		}
	}

}

func getOriginalCurrency() (string, error) {
	var userInput string
	fmt.Println("Выбрите валюту: (USD/EUR/RUB)")
	fmt.Scan(&userInput)
	if userInput != "USD" && userInput != "EUR" && userInput != "RUB" {
		return "", errors.New("Некорректная валюта")
	}
	return userInput, nil
}
func getSum() (float64, error) {
	var sum float64
	fmt.Println("Введите сумму: ")
	fmt.Scan(&sum)
	if sum <= 0 {
		return 0, errors.New("Некорректная сумма")
	}
	return sum, nil
}
func getTargetCurrency(origCur string) (string, error) {
	var targetCurrency string
	switch {
	case origCur == "USD":
		fmt.Println("В какую валюту хотите конвертировать? (EUR/RUB)")
	case origCur == "EUR":
		fmt.Println("В какую валюту хотите конвертировать? (USD/RUB)")
	case origCur == "RUB":
		fmt.Println("В какую валюту хотите конвертировать? (USD/EUR)")
	}
	fmt.Scan(&targetCurrency)
	if targetCurrency != "USD" && targetCurrency != "EUR" && targetCurrency != "RUB" {
		return "", errors.New("Некорректная валюта")
	} else if targetCurrency == "USD" && origCur == "USD" {
		return "", errors.New("Невозможно конвертировать USD в USD")
	} else if targetCurrency == "RUB" && origCur == "RUB" {
		return "", errors.New("Невозможно конвертировать RUB в RUB")
	} else if targetCurrency == "EUR" && origCur == "EUR" {
		return "", errors.New("Невозможно конвертировать EUR в EUR")
	}
	return targetCurrency, nil
}

func convertCurrency(sum float64, originalCurrency string, targetCurrency string) float64 {
	var convertedCurrency float64
	switch {
	case originalCurrency == "USD" && targetCurrency == "EUR":
		convertedCurrency = sum * usdEur
		return convertedCurrency
	case originalCurrency == "USD" && targetCurrency == "RUB":
		convertedCurrency = sum * usdRub
		return convertedCurrency
	case originalCurrency == "EUR" && targetCurrency == "RUB":
		convertedCurrency = sum * eurRub
		return convertedCurrency
	case originalCurrency == "EUR" && targetCurrency == "USD":
		convertedCurrency = sum * eurUsd
		return convertedCurrency
	case originalCurrency == "RUB" && targetCurrency == "USD":
		convertedCurrency = sum * rubUsd
		return convertedCurrency
	case originalCurrency == "RUB" && targetCurrency == "EUR":
		convertedCurrency = sum * rubEur
		return convertedCurrency
	}
	return convertedCurrency
}
func wantToRestart() bool {
	var restart string
	fmt.Println("Хотите повторить вычисление? (y/n)")
	fmt.Scan(&restart)
	if restart == "y" {
		return true
	} else {
		return false
	}

}
