package main

import "fmt"

func minSwaps(s string) int {
	zero := 0
	one := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	if zero-one > 1 || one-zero > 1 {
		return -1
	}
	if zero == one {
		return min(calcNumSwap(s, '0'), calcNumSwap(s, '1'))
	}
	if zero > one {
		return calcNumSwap(s, '0')
	}
	return calcNumSwap(s, '1')
}

func calcNumSwap(s string, c byte) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			res++
		}
		c = c ^ 1
	}
	return res / 2
}

func main() {
	var s string = "111000"
	fmt.Println(minSwaps(s))
}
