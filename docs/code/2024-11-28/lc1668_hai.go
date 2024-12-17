package main

import "fmt"

func maxRepeating(sequence string, word string) int {
	var cnt int
    target:=""
    for i := 1; i <= len(sequence)/len(word);i++ {
        target+=word
        for j:=0; j< len(sequence);j++ {
            if len(sequence) >= j + len(target) && sequence[j:j+len(target)] == target {
                cnt=i
            }
        }
    }
    return cnt
}

func main() {
	sequence := "ababc"
	word := "ab"
	fmt.Println(maxRepeating(sequence, word)) // 2
}
