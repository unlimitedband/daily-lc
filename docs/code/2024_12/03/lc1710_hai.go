package main

import "fmt"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var ans int
	var n int = len(boxTypes)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if boxTypes[i][1] < boxTypes[j][1] {
				boxTypes[i], boxTypes[j] = boxTypes[j], boxTypes[i]
			}
		}
	}
	for i := 0; i < n; i++ {
		if truckSize > boxTypes[i][0] {
			ans += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			ans += truckSize * boxTypes[i][1]
			break
		}
	}
	return ans
}

func main() {
	boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	truckSize := 4
	fmt.Println(maximumUnits(boxTypes, truckSize)) // 8
}
