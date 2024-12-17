package main

import "sort"

func matrixSum(nums [][]int) int {
	n, m := len(nums), len(nums[0])
	res := 0

	for i := 0; i < n; i++ {
		sort.Ints(nums[i])
	}

	for j := 0; j < m; j++ {
		maxSum := 0
		for i := 0; i < n; i++ {
			maxSum = max(maxSum, nums[i][j])
		}
		res += maxSum
	}
	return res
}

func main() {
	nums := [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}
	println(matrixSum(nums)) // 15
}
