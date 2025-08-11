package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		operation := ImportOperation()
		numbers := ImportNumbers()
		itog := RealOperation(operation, numbers)
		fmt.Printf("Итог вашей операции: %v \n", itog)
		fmt.Println("Повторить?: ")
		var ans string
		fmt.Scan(&ans)
		answer := CheckANS(ans)
		if !answer {
			break
		}
	}
}

func ImportOperation() string {
	var operation string
	fmt.Println("Введите операцию: ")
	fmt.Scan(&operation)
	return operation
}
func ImportNumbers() []int {
	var numbers string
	fmt.Println("Введите числа (через запятую): ")
	fmt.Scan(&numbers)
	num := strings.Split(numbers, ",")
	var nums []int
	for _, s := range num {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	return nums
}

func AWG(numbers []int) float64 {
	return float64(SUM(numbers)) / float64(len(numbers))
}

func SUM(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func MED(numbers []int) float64 {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return float64(sorted[n/2-1]+sorted[n/2]) / 2
}

func RealOperation(operation string, numbers []int) interface{} {
	switch {
	case operation == "AWG":
		return AWG(numbers)
	case operation == "SUM":
		return SUM(numbers)
	case operation == "MED":
		return MED(numbers)

	}
	return ""
}

func CheckANS(usans string) bool {
	return usans == "y" || usans == "Y"
}
