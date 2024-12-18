package main

import "fmt"

func beautifulSplits(nums []int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] == nums[j] {
				f[i][j] = f[i+1][j+1] + 1
			}
		}
	}
    var res int
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			a := i <= j-i && f[0][i] >= i
			b := j-i <= n-j && f[i][j] >= j-i
			if a || b {
				res++
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 2, 1}
	fmt.Print(beautifulSplits(nums))
}
