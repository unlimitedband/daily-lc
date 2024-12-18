package main

import "fmt"

func numsSameConsecDiff(n int, k int) []int {
	var ret []int
	if n == 1 {
		for num := 1; num <= 9; num++ {
			if num-k >= 0 || num+k <= 9 {
				ret = append(ret, num)
			}
		}
		return ret
	} else {
		ret = numsSameConsecDiff(n-1, k)
		var tmp []int
		for _, num := range ret {
			lastDigit := num % 10
			if lastDigit-k >= 0 {
				tmp = append(tmp, num*10+lastDigit-k)
			}
			if lastDigit+k <= 9 && k != 0 {
				tmp = append(tmp, num*10+lastDigit+k)
			}
		}
		return tmp
	}

}

func main() {
	fmt.Println(numsSameConsecDiff(2, 0))
}
