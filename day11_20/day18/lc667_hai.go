package main

import "fmt"

func constructArray(n int, k int) []int {
	ret := make([]int, n)
	be := 1
	en := n + 1
	add := 0
	for i := 0; i < k-1; i++ {
		ret[i] = be
		be -= add
		be = en - be
		add = 1 - add
	}
	//fmt.Println(be)
	if be > n/2 {
		for i := k - 1; i < n; i++ {
			ret[i] = be
			be--
		}
	} else {
		for i := k - 1; i < n; i++ {
			ret[i] = be
			be++
		}
	}
	return ret
}

func main() {
	n := 5
	k := 4
	fmt.Println(constructArray(n, k))
}
