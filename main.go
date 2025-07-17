package main

import (
	"errors"
	"fmt"
)

const UsdToEur = 0.85
const UsdToRub = 78.6

func main() {
	for {
		fmt.Println("-Конвертер валют-")
		var orgValue string
		var err error
		var tarValue string
		var errr error
		var sum float64
		var erro error
		for {
			orgValue, err = getUserOValue()
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				break
			}
		}
		for {
			sum, errr = getUserSum()
			if errr != nil {
				fmt.Println(errr)
				continue
			} else {
				break
			}
		}
		for {
			tarValue, erro = checkInputValue(orgValue)
			if erro != nil {
				fmt.Println(erro)
				continue
			} else {
				break
			}
		}
		result := converter(sum, orgValue, tarValue)
		fmt.Printf("Итоговый результат перевода валют: %.2f %s", result, tarValue)
		var userAns string
		fmt.Println("Повторить запрос(y/n) ")
		fmt.Scan(&userAns)
		ans := checkAns(userAns)
		if !ans {
			break
		}
	}
}

func getUserOValue() (string, error) {
	var orgValue string
	fmt.Print("Введите исходную валюту(usd/eur/rub): ")
	fmt.Scan(&orgValue)
	if orgValue != "usd" && orgValue != "eur" && orgValue != "rub" {
		return "", errors.New("вы ввели не ту валюту, повторите попытку")
	}
	return orgValue, nil
}

func getUserSum() (float64, error) {
	var sum float64
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&sum)
	if sum <= 0 {
		return 0, errors.New("Неправильная сумма ввода")
	}
	return sum, nil
}

func checkInputValue(orgValue string) (string, error) {
	var tarValue string
	switch {
	case "usd" == orgValue:
		fmt.Print("Введите целевую валюту (eur/rub): ")
		fmt.Scan(&tarValue)
		if "eur" != tarValue && "rub" != tarValue {
			return "", errors.New("вы ввели не ту валюту, повторите попытку")
		}
	case "eur" == orgValue:
		fmt.Print("Введите целевую валюту (usd/rub): ")
		fmt.Scan(&tarValue)
		if "usd" != tarValue && "rub" != tarValue {
			return "", errors.New("вы ввели не ту валюту, повторите попытку")
		}
	case "rub" == orgValue:
		fmt.Print("Введите целевую валюту (usd/eur): ")
		fmt.Scan(&tarValue)
		if "usd" != tarValue && "eur" != tarValue {
			return "", errors.New("вы ввели не ту валюту, повторите попытку")
		}
	default:
		return tarValue, nil
	}
	return tarValue, nil
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

func checkAns(ans string) bool {
	if ans == "y" || ans == "Y" {
		return true
	} else {
		return false
	}
}
