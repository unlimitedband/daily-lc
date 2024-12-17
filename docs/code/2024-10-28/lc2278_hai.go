package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	var cnt int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			cnt++
		}
	}
	return cnt * 100 / len(s)
}

func main() {
	s := "aaabbbccc"
	var letter byte = 'a'
	fmt.Println(percentageLetter(s, letter))
}
