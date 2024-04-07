package main

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	starts := make([][]int, n)
	for i := 0; i < n; i++ {
		starts[i] = make([]int, 2)
		starts[i][0] = intervals[i][0]
		starts[i][1] = i
	}
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})
	var res []int
	for _, interval := range intervals {
		l, r, end := 0, n-1, interval[1]
		for l < r {
			mid := (l + r) / 2
			if starts[mid][0] >= end {
				r = mid
			} else {
				l = mid + 1
			}
		}
		val := -1
		if starts[l][0] >= end {
			val = starts[l][1]
		}
		res = append(res, val)
	}
	return res
}

func main() {
	intervals := [][]int{{3, 4}, {2, 3}, {1, 2}}
	fmt.Println(findRightInterval(intervals))
}
