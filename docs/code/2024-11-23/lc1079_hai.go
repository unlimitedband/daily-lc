package main

func numTilePossibilities(tiles string) int {
	cnt := make([]int, 26)
	for _, v := range tiles {
		cnt[v-'A']++
	}
	return dfs(cnt)
}

func dfs(cnt []int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if cnt[i] == 0 {
			continue
		}
		sum++
		cnt[i]--
		sum += dfs(cnt)
		cnt[i]++
	}
	return sum
}

func main() {
	tiles := "AAB"
	println(numTilePossibilities(tiles)) // 8
}
