package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	var box [46]int
	var ans int
	for i := lowLimit; i <= highLimit; i++ {
		var sum int
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		box[sum]++
		if box[sum] > ans {
			ans = box[sum]
		}
	}
	return ans
}

func main() {
	lowLimit := 1
	highLimit := 10
	fmt.Println(countBalls(lowLimit, highLimit)) // 2
}
