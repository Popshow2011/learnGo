package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.88
	const USD_TO_RUB = 82.50
	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR
	fmt.Println(EUR_TO_RUB)

}
