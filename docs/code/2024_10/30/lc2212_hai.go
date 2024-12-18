package main

import "fmt"

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	n := len(aliceArrows)
	maxTurn := 0
	maxPoint := -1
	for mask := 1; mask < 1<<n; mask++ {
		cnt := 0
		points := 0
		for i, alice := range aliceArrows {
			if (mask>>i)&1 == 1 {
				cnt += alice + 1
				points += i
			}
		}
		if cnt <= numArrows && maxPoint < points {
			maxTurn = mask
			maxPoint = points
		}
	}
	ans := make([]int, n)
	for i, alice := range aliceArrows {
		if (maxTurn>>i)&1 == 1 {
			ans[i] = alice + 1
			numArrows -= ans[i]
		}
	}
	ans[0] += numArrows
	return ans
}

func main() {
	numArrows := 5
	aliceArrows := []int{1, 2, 3}
	fmt.Println(maximumBobPoints(numArrows, aliceArrows))
}
