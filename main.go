package main

import (
	"fmt"
)

func main() {
	const UsdToEur = 0.85
	const UsdToRub = 78.6
	EurToRub := UsdToRub / UsdToEur
	fmt.Print(EurToRub)
}
func getUserInput() (string, string) {
	var orgValue string
	var tarValue string
	fmt.Scan(&orgValue)
	fmt.Scan(&tarValue)
	return orgValue, tarValue
}

func convertate(orgValue string, tarValue string, sum int) int {
	return 0
}
