package main

import "fmt"

func minMoves(target int, maxDoubles int) int {
	if maxDoubles == 0 {
		return target - 1
	}
	if target == 1 {
		return 0
	}
	if target%2 == 0 {
		return minMoves(target/2, maxDoubles-1) + 1
	} else {
		return minMoves(target-1, maxDoubles) + 1
	}
}

func main() {
	target := 10
	maxDoubles := 4
	fmt.Println(minMoves(target, maxDoubles))
}
