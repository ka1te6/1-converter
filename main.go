package main

import "fmt"

func main() {
	const UsdToEur = 0.85
	const UsdToRub = 78.6
	EurToRub := UsdToRub / UsdToEur
	fmt.Print(EurToRub)
}
