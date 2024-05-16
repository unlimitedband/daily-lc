package main

import "fmt"

func differenceOfSums(n int, m int) int {
	if m == 0 {
		return n
	}
	return n*(n+1)/2 - m*(n/m)*(n/m+1)
}

func main() {
	n, m := 10, 3
	fmt.Println(differenceOfSums(n, m)) // 19
}
