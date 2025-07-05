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
func getUserInput() int {
	var sum int
	fmt.Scan(&sum)
	return sum
}

func convertate(sum int, orgValue string, tarValue string) int {
	return 0
}
