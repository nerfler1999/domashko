package main

import (
	"errors"
	"fmt"
)

func main() {
	constMap := map[string]float64{
		"USD-EUR": 0.85,
		"USD-RUB": 75.14,
		"EUR-RUB": 83.06,
		"EUR-USD": 1.17,
		"RUB-USD": 0.013,
		"RUB-EUR": 0.011,
	}
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
		total := convertCurrency(sum, cur, targ, &constMap)
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

func convertCurrency(sum float64, originalCurrency string, targetCurrency string, constMap *map[string]float64) float64 {

	key := originalCurrency + "-" + targetCurrency
	return sum * (*constMap)[key]
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
