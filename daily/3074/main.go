package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hallo!")
	apple1 := []int{1, 3, 2}
	capacity1 := []int{4, 3, 1, 5, 2}
	apple2 := []int{5, 5, 5}
	capacity2 := []int{2, 4, 2, 7}
	fmt.Println(minimumBoxes(apple1, capacity1))
	fmt.Println(minimumBoxes(apple2, capacity2))
}

// Optimization was using the first entry of apples
// as the place to store the apple count and saving one variable
func minimumBoxes(apple []int, capacity []int) int {
	boxCount := 0
	for i := 1; i < len(apple); i++ {
		apple[0] += apple[i]
	}
	capacities := len(capacity)
	sort.Ints(capacity)
	for i := range capacities {
		boxCount += 1
		apple[0] -= capacity[capacities-1-i]
		if apple[0] <= 0 {
			break
		}
	}
	return boxCount
}
