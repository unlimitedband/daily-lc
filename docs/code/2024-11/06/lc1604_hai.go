package main

import (
	"fmt"
	"sort"
)

func alertNames(keyName []string, keyTime []string) []string {
	timeMap := make(map[string][]int)
	for i := 0; i < len(keyName); i++ {
		hour := int(keyTime[i][0]-'0')*10 + int(keyTime[i][1]-'0')
		minute := int(keyTime[i][3]-'0')*10 + int(keyTime[i][4]-'0')
		time := hour*60 + minute
		timeMap[keyName[i]] = append(timeMap[keyName[i]], time)
	}
	for _, v := range timeMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}
	res := []string{}
	for k, v := range timeMap {
		for i := 2; i < len(v); i++ {
			if v[i]-v[i-2] <= 60 {
				res = append(res, k)
				break
			}
		}
	}

	sort.Strings(res)
	return res

}

func main() {
	keyName := []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}
	keyTime := []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}
	fmt.Println(alertNames(keyName, keyTime)) // ["daniel"]
}
