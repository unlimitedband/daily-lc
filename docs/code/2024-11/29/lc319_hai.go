package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	n := 3
	fmt.Println(bulbSwitch(n)) // 1
}
