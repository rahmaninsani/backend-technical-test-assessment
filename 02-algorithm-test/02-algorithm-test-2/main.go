package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProf := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice

		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if profit > maxProf {
			maxProf = profit
		}
	}

	return maxProf
}

func main() {
	priceList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{7, 9},
	}

	for _, prices := range priceList {
		result := maxProfit(prices)
		fmt.Printf("Hasil %v: %d\n", prices, result)
	}
}
