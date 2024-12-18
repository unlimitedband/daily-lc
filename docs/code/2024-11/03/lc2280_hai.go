package main

import (
	"fmt"
	"sort"
)

func minimumLines(stockPrices [][]int) int {
	ret := 1
	if len(stockPrices) <= 2 {
		return len(stockPrices) - 1
	}
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})
	slopeA := stockPrices[1][1] - stockPrices[0][1]
	slopeB := stockPrices[1][0] - stockPrices[0][0]
	for i := 2; i < len(stockPrices); i++ {
		slopeA1 := stockPrices[i][1] - stockPrices[i-1][1]
		slopeB1 := stockPrices[i][0] - stockPrices[i-1][0]
		if slopeA*slopeB1 != slopeB*slopeA1 {
			ret++
			slopeA = slopeA1
			slopeB = slopeB1
		}
	}
	return ret
}

func main() {
	fmt.Println(minimumLines([][]int{{72, 98}, {62, 27}, {32, 7}, {71, 4}, {25, 19}, {91, 30}, {52, 73}, {10, 9}, {99, 71}, {47, 22}, {19, 30}, {80, 63}, {18, 15}, {48, 17}, {77, 16}, {46, 27}, {66, 87}, {55, 84}, {65, 38}, {30, 9}, {50, 42}, {100, 60}, {75, 73}, {98, 53}, {22, 80}, {41, 61}, {37, 47}, {95, 8}, {51, 81}, {78, 79}, {57, 95}}))

}
