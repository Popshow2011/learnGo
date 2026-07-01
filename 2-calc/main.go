package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		userChoice := 0
		fmt.Println("Выберите нужный пункт меню")
		fmt.Println("1. AVG")
		fmt.Println("2. SUM")
		fmt.Println("3. MED")
		fmt.Println("4. Выход")
		fmt.Scan(&userChoice)
		switch userChoice {
		case 1:
			avg := getAvg()
			fmt.Println(avg)
		case 2:
			sum := getSum()
			fmt.Println(sum)
		case 3:
			med := getMedian()
			fmt.Println(med)
		default:
			return
		}
	}

}

func convertStrToNum(str string) []float64 {
	parts := strings.Split(str, ",")
	arr := make([]float64, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			panic("ERROR: ошибка при конвертации числа")
		}
		arr = append(arr, num)
	}
	return arr
}

func getAvg() float64 {
	var sum float64
	prompt := ""
	fmt.Println("Введите цифры через запятую")
	fmt.Scan(&prompt)

	arr := convertStrToNum(prompt)

	for _, num := range arr {
		sum += num
	}
	avg := sum / float64(len(arr))
	return avg

}

func getSum() float64 {
	var sum float64
	prompt := ""
	fmt.Println("Введите цифры через запятую")
	fmt.Scan(&prompt)
	arr := convertStrToNum(prompt)

	for _, num := range arr {
		sum += num
	}
	return sum

}

func getMedian() float64 {
	prompt := ""
	fmt.Println("Введите цифры через запятую")
	fmt.Scan(&prompt)
	numbers := convertStrToNum(prompt)

	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)

	sort.Float64s(sorted)

	n := len(sorted)
	if n == 0 {
		return 0
	}

	if n%2 != 0 {
		return sorted[n/2]
	} else {
		mid1 := sorted[n/2-1]
		mid2 := sorted[n/2]
		return (mid1 + mid2) / 2.0
	}
}
