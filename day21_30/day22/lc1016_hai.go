package main

import "fmt"

func queryString(s string, n int) bool {
	checkNum := map[int]bool{}
	for i := 0; i < len(s); i++ {
		num := 0
		for j := i; j < len(s); j++ {
			num = num*2 + int(s[j]-'0')
			if num > n {
				break
			}
			checkNum[num] = true
		}
	}
	for i := 1; i <= n; i++ {
		if !checkNum[i] {
			return false
		}
	}

	return true
}

func main() {
	var s string = "0110"
	var n int = 3
	fmt.Println(queryString(s, n))
}
