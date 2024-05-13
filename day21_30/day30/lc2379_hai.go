package main

func minimumRecolors(blocks string, k int) int {
	result, N, cnt := 100, len(blocks), 0
	for l, r := 0, 0; r < N; r++ {
		if blocks[r] == 'W' {
			cnt++
		}
		for r-l+1 == k {
			if cnt < result {
				result = cnt
			}
			if blocks[l] == 'W' {
				cnt--
			}
			l++
		}
	}
	return result
}

func main() {
	blocks := "WBBWWBBWBW"
	k := 7
	println(minimumRecolors(blocks, k)) // 3
}
