package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	var n int = len(nums)
	var ans int = 0
	preSum := make([]int, n+1)
	firstSum := make([]int, n+1)
	secondSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] += preSum[i-1] + nums[i-1]
		if i >= firstLen {
			firstSum[i] = max(firstSum[i-1], preSum[i]-preSum[i-firstLen])
		}
		if i >= secondLen {
			secondSum[i] = max(secondSum[i-1], preSum[i]-preSum[i-secondLen])
		}
	}
	for i := firstLen + secondLen; i <= n; i++ {
		first := preSum[i] - preSum[i-secondLen] + firstSum[i-secondLen]
		second := preSum[i] - preSum[i-firstLen] + secondSum[i-firstLen]
		ans = max(ans, max(first, second))
	}
	return ans
}

func main() {
	nums := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
	firstLen := 1
	secondLen := 2
	fmt.Println(maxSumTwoNoOverlap(nums, firstLen, secondLen)) // 20
}
