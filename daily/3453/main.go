package main

import "fmt"

func main() {
	testCase1 := [][]int{
		{0, 0, 1},
		{2, 2, 1},
	}
	fmt.Println(separateSquares(testCase1))
}

func separateSquares(squares [][]int) float64 {
	var totalArea float64
	var low float64 = 10e9
	var high float64 = -10e9

	for _, s := range squares {
		y, l := float64(s[1]), float64(s[2])
		totalArea += l * l
		if y < low {
			low = y
		}
		if y+l > high {
			high = y + l
		}
	}

	halfArea := totalArea / 2.0
	for range(47) {
		mid := low + (high-low)/2
		if getAreaBelow(mid, squares) < halfArea {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func getAreaBelow(mid float64, squares [][]int) float64 {
	var areaBelow float64 = 0
	for _, s := range squares {
		y := float64(s[1])
		l := float64(s[2])
		top := y + l

		if mid <= y {
			continue
		} else if mid >= top {
			areaBelow += l * l
		} else {
			areaBelow += l * (mid - y)
		}
	}
	return areaBelow
}
