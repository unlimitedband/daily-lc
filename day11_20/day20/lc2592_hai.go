package main

import (
	"fmt"
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	var maxR int = 0
	for i := 0; i < len(ranks); i++ {
		maxR = max(maxR, ranks[i])
	}
	var maxTimeToRepairCar int64 = int64(maxR) * int64(cars) * int64(cars)
	return findMinTime(0, maxTimeToRepairCar, ranks, cars)
}

func findMinTime(l int64, r int64, ranks []int, cars int) int64 {
	var ret int64 = -1
	for l <= r {
		m := (l + r) / 2
		if checkTime(ranks, cars, m) {
			ret = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ret
}

func checkTime(ranks []int, cars int, time int64) bool {
	cntCar := 0
	for i := 0; i < len(ranks); i++ {
		cntCar += int(math.Sqrt(float64(time) / float64(ranks[i])))
	}
	return cntCar >= cars
}

func main() {
	ranks := []int{4, 6, 4, 5, 4, 3}
	cars := 1441
	fmt.Println(repairCars(ranks, cars))
}
