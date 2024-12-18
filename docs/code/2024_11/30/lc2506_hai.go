package main

import "fmt"

func similarPairs(words []string) int {

	cnt := make(map[[26]int]int)

	for _, word := range words {
		var set [26]int
		for _, c := range word {
			set[c-'a'] = 1
		}
		cnt[set]++
	}

	res := 0
	for _, v := range cnt {
		if v > 1 {
			res += v * (v - 1) / 2
		}
	}

	return res
}

func main() {
	words := []string{"aba", "aabb", "abcd", "bac", "aabc"}
	fmt.Println(similarPairs(words)) // 4
}
