package main

import "fmt"

func main() {
	customers1 := "YYNY"
	customers2 := "NNNNN"
	customers3 := "YYYY"
	customers4 := "YNYY"

	fmt.Println(bestClosingTime(customers1))
	fmt.Println()
	fmt.Println(bestClosingTime(customers2))
	fmt.Println()
	fmt.Println(bestClosingTime(customers3))
	fmt.Println()
	fmt.Println(bestClosingTime(customers4))
}

func bestClosingTime(customers string) int {
	n := len(customers)
	bestTime := n
	penality := 0

	minimalPenality := penality

	for i := range n {
		if customers[n-1-i] == 'N' {
			penality -= 1
		} else {
			penality += 1
		}

		if penality <= minimalPenality {
			bestTime = n - i - 1
			minimalPenality = penality
		}
	}
	return bestTime
}
