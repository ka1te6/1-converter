package main

import "fmt"

func main() {
	const UE = 0.85
	const UR = 78.6
	ER := (1 / UE) * UR
	fmt.Print(ER)
}
