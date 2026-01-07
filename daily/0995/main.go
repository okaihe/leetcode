package main

import "fmt"

func main() {
	strs1 := []string{"ca", "bb", "ac"}
	strs2 := []string{"xc", "yb", "za"}
	strs3 := []string{"zyx", "wvu", "tsr"}
	fmt.Println(minDeletionSize([]string{}))
	fmt.Println(minDeletionSize(strs1))
	fmt.Println(minDeletionSize(strs2))
	fmt.Println(minDeletionSize(strs3))
}

func minDeletionSize(strs []string) int {
	n := len(strs)
	if n == 0 {
		return 0
	}

	stringSize := len(strs[0])
	deletionCount := 0
	isSorted := make([]bool, n-1)

outer:
	for i := range stringSize {
		for j := 0; j < n-1; j++ {
			if !isSorted[j] && strs[j][i] > strs[j+1][i] {
				deletionCount++
				continue outer
			}
		}
		for j := 0; j < n-1; j++ {
			if strs[j][i] < strs[j+1][i] {
				isSorted[j] = true
			}
		}
	}
	return deletionCount
}
