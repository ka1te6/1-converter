package main

import (
	"errors"
	"fmt"
)

const UsdToEur = 0.85
const UsdToRub = 78.6

func main() {
	fmt.Println("-Конвертер валют-")
	sum, orgValue, tarValue, err := getUserInput()
	if err != nil {
		fmt.Println(err)
	}
	result := converter(sum, orgValue, tarValue)
	fmt.Printf("Итоговый результат перевода валют %.2f", result)
	var userAns string
	fmt.Println("Повторить запрос(y/n)")
	fmt.Scan(&userAns)

}
func getUserInput() (float64, string, string, error) {
	var orgValue string
	var tarValue string
	var sum float64
	fmt.Print("Введите исходную валюту(usd/eur/rub): ")
	fmt.Scan(&orgValue)
	if orgValue != "usd" || orgValue != "eur" || orgValue != "rub" {
		return 0, "", "", errors.New("вы ввели не ту валюту, повторите попытку")
	}
	checkInput(orgValue)
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&sum)
	if sum <= 0 {
		return 0, "", "", errors.New("Неправильная сумма ввода")
	}
	return sum, orgValue, tarValue, nil
}

func checkInput(orgValue string) (string, string, error) {
	var tarValue string
	switch {
	case "usd" == orgValue:
		fmt.Print("Введите целевую валюту (eur/rub): ")
		fmt.Scan(&tarValue)
		if "eur" != tarValue || "rub" != tarValue {
			return "", "", errors.New("вы ввели не ту валюту, повторите попытку")
		}
	case "eur" == orgValue:
		fmt.Print("Введите целевую валюту (usd/rub): ")
		fmt.Scan(&tarValue)
		if "usd" != tarValue || "rub" != tarValue {
			return "", "", errors.New("вы ввели не ту валюту, повторите попытку")
		}
	case "rub" == orgValue:
		fmt.Print("Введите целевую валюту (usd/eur): ")
		fmt.Scan(&tarValue)
		if "usd" != tarValue || "eur" != tarValue {
			return "", "", errors.New("вы ввели не ту валюту, повторите попытку")
		}
	default:
		return orgValue, tarValue, nil
	}
	return orgValue, tarValue, nil
}

func converter(sum float64, orgValue string, tarValue string) float64 {
	switch {
	case orgValue == "usd" && tarValue == "eur":
		return UsdToEur * sum
	case orgValue == "usd" && tarValue == "rub":
		return UsdToRub * sum
	case orgValue == "eur" && tarValue == "usd":
		return (1 / UsdToEur) * sum
	case orgValue == "eur" && tarValue == "rub":
		return (UsdToRub / UsdToEur) * sum
	case orgValue == "rub" && tarValue == "usd":
		return (1 / UsdToRub) * sum
	case orgValue == "rub" && tarValue == "eur":
		return (UsdToEur / UsdToRub) * sum
	default:
		return 0
	}
}

func checkAns(string) bool {
	var ans string
	if ans == "y" || ans == "Y" {
		return true
	} else {
		return false
	}
}
