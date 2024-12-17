package main

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	mod := grid[0][0] % x
	fmt.Println(mod)
	arr := []int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j]%x != mod {
				return -1
			}
			arr = append(arr, (grid[i][j]-mod)/x)
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	res := 0
	for i := 0; i < len(arr); i++ {
		res += int(math.Abs(float64(arr[i] - arr[len(arr)/2])))
	}
	return res
}

func main() {
	// grid := [][]int{{2, 4}, {6, 8}}
	grid := [][]int{{146}}
	x := 86
	fmt.Println(minOperations(grid, x)) // 4
}
