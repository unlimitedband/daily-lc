package main

import (
	"fmt"
	"math"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	date01, _ := time.Parse("2006-01-02", date1)
	date02, _ := time.Parse("2006-01-02", date2)
	difference := date02.Sub(date01)
	return int(math.Abs(difference.Hours() / 24))
}

func main() {
	date1 := "2019-06-29"
	date2 := "2019-06-30"
	fmt.Println(daysBetweenDates(date1, date2)) // 1

	date1 = "2020-01-15"
	date2 = "2019-12-31"
	fmt.Println(daysBetweenDates(date1, date2)) // 15
}
