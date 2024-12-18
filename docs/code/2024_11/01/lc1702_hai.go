package main

import (
	"fmt"
)

func maximumBinaryString(binary string) string {
	firstZeroPos := -1
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			firstZeroPos = i
			break
		}
	}
	if firstZeroPos == -1 {
		return binary
	}
	for _, digit := range binary[firstZeroPos+1:] {
		if digit == '0' {
			firstZeroPos++
		}
	}
	ret := []byte(binary)
	for i := range ret {
		ret[i] = '1'
	}
	ret[firstZeroPos] = '0'
	return string(ret)
}

func main() {
	fmt.Println(maximumBinaryString("000110"))
}
