package main

import "fmt"

func majorityElement(nums []int) int {
	majority, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
			count++
		} else if majority == nums[i] {
			count++
		} else {
			count--
		}
	}
	return majority
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums)) // 2
}
