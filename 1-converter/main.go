package main

import (
	"fmt"
	"strings"
)

const USD_TO_EUR = 0.88
const USD_TO_RUB = 82.50
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

func main() {
	initialCurr, value, targetCurr := getUserInput()
	result := converter(value, initialCurr, targetCurr)
	fmt.Println(result)

}

func availableCurrency(curr string) string {
	switch strings.ToUpper(curr) {
	case "USD":
		return "EUR RUB"
	case "EUR":
		return "USD RUB"
	default:
		return "USD EUR"
	}
}

func converter(value float64, initial, target string) float64 {
	initial = strings.ToUpper(initial)
	target = strings.ToUpper(target)

	if initial == target {
		return value
	}

	var inUSD float64
	switch initial {
	case "USD":
		inUSD = value
	case "EUR":
		inUSD = value / USD_TO_EUR
	case "RUB":
		inUSD = value / USD_TO_RUB
	}

	switch target {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * USD_TO_EUR
	case "RUB":
		return inUSD * USD_TO_RUB
	}

	return 0
}

func getUserInput() (string, float64, string) {
	var initialCurr string
	var targetCurr string
	var value float64

	for initialCurr == "" {
		fmt.Println("Введите наименование исходной валюты: USD RUB EUR")
		_, err := fmt.Scan(&initialCurr)
		if err != nil {
			fmt.Println(err.Error())
		}
		isInitial := strings.Contains("EUR RUB USD", strings.ToUpper(initialCurr))
		if !isInitial {
			fmt.Printf("Вы ввели недопустимую валюту - %s, попробуйте еще раз\n", initialCurr)
			initialCurr = ""
			continue
		}
	}

	for value == 0 {
		fmt.Println("Введите колличество валюты")
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println(err.Error())
		}
		if value <= 0 {
			fmt.Println("Число должно быть больше 0")
			value = 0
			continue
		}
	}
	for targetCurr == "" {
		availableCurr := availableCurrency(initialCurr)
		fmt.Printf("Введите наименование исходной валюты: %s \n", availableCurr)
		_, err := fmt.Scan(&targetCurr)
		if err != nil {
			fmt.Println(err.Error())
		}
		isTargetCurr := strings.Contains(strings.ToUpper(availableCurr), strings.ToUpper(targetCurr))
		if !isTargetCurr {
			fmt.Printf("Вы ввели недопустимую валюту - %s, попробуйте еще раз\n", availableCurr)
			targetCurr = ""
			continue
		}
	}
	return initialCurr, value, targetCurr
}
