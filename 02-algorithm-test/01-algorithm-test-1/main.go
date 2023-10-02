package main

import (
	"fmt"
	"math"
)

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProductOfThree(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num > max1 {
			max3, max2, max1 = max2, max1, num
		} else if num > max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min2, min1 = min1, num
		} else if num < min2 {
			min2 = num
		}
	}

	maxPositive := max1 * max2 * max3
	maxNegative := min1 * min2 * max1

	return findMax(maxPositive, maxNegative)
}

func main() {
	nums := [][]int{
		{1, 2, 3},
		{-10, -10, 1, 3, 2},
		{1, 2, 3, 4, 5},
		{0, 0, 0, 0, 0},
		{-5, 0, 1, 2, 3},
	}

	for _, num := range nums {
		result := maxProductOfThree(num)
		fmt.Printf("Hasil %v: %d\n", num, result)
	}
}
