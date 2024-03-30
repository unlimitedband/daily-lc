/*
Yêu cầu:
Có 2 action với một số N cho trước là cộng 2^x hoặc trừ 2^x vào số N.
Tính số action để số N bằng 0.


Ý tưởng:
Biểu diễn số N sang hệ nhị phân
-> Chi phí đổi 1 bit mang giá trị 1 <= 1 (trừ 2 ^ x với x là vị trí của bit từ phải qua trái)
VD N = 100 -> trừ 2 ^ 2
-> Chi phí đổi dãy n bit liên tiếp <= 2 (cộng z ^ x1 với x1 là vị trí đầu của dãy bit, từ 2 ^ x2 với x2 là vị trí cuối tính từ phải qua trái)
VD: N = 11100 -> cộng 2 ^ 2 N sẽ thành 100000 -> trừ 2 ^ 5

*/

package main

import (
	"fmt"
)

func minOperations(n int) int {
	ans := 0
	numberOfConsecutiveBit1 := 0
	index := 0
	minBit1Pos := -1
	maxBit1Pos := -1

	for n != 0 {
		if (n>>index)&1 == 1 {
			if numberOfConsecutiveBit1 == 0 {
				minBit1Pos = index
			}
			maxBit1Pos = index
			numberOfConsecutiveBit1++
		} else {
			if maxBit1Pos == -1 || numberOfConsecutiveBit1 == 0 {
				index++
				continue
			}
			if numberOfConsecutiveBit1 == 1 {
				n = n - 1<<minBit1Pos
				maxBit1Pos = -1
				minBit1Pos = -1
				numberOfConsecutiveBit1 = 0
				ans++
			} else {
				n = n + 1<<minBit1Pos
				minBit1Pos = index
				numberOfConsecutiveBit1 = 1
				maxBit1Pos = index
				ans++
			}
		}
		index++
	}
	return ans
}

func main() {
	var n int = 39
	fmt.Println(minOperations(n))
}
