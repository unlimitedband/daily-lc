package main

import "fmt"

const mod = 1e9 + 7

func pow2(n int) int {
	if n == 0 {
		return 1
	}
	a := pow2(n / 2)
	if n%2 == 0 {
		return a * a % mod
	} else {
		return a * a % mod * 2 % mod
	}
}

func productQueries(n int, queries [][]int) []int {
	var bit []int
	bit = append(bit, 0)
	for i := 0; n > 0; i++ {
		if n%2 == 1 {
			bit = append(bit, i)
		}
		n /= 2
	}
	for i := 1; i < len(bit); i++ {
		bit[i] = bit[i-1] + bit[i]
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]

		ans[i] = pow2(bit[r+1]-bit[l]) % mod
	}
	return ans
}

func main() {
	n := 15
	queries := [][]int{{0, 0}, {2, 2}, {0, 3}}
	fmt.Println(productQueries(n, queries)) // [2,4,64]
}
