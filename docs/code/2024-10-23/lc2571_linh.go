package main

import "fmt"

func minOperations(n int) int {
	if n == 1 {
		return 1
	}

	m := 1
	for m < n {
		m <<= 1
	}

	if m == n {
		return 1
	}
	if m-n > n-(m>>1) {
		return 1 + minOperations(n-(m>>1))
	}
	return 1 + minOperations(m-n)
}

func main() {
	n := 39
	fmt.Println(minOperations(n))
}
