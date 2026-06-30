package main

import "fmt"

const USD_TO_EUR = 0.88
const USD_TO_RUB = 82.50
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

func main() {
	prompt, err := userPrompt()
	if err != nil {
		fmt.Println(err.Error())
	}
	converter(prompt, "USD", "EUR")
}

func converter(value int, initial, target string) {}

func userPrompt() (int, error) {
	var p int = 0
	fmt.Scan(&p)

	if p == 0 {
		return 0, fmt.Errorf("[ERROR]: Пользователь ни чего не ввел")
	}
	return p, nil
}
