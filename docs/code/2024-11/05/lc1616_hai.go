package main

import "fmt"

func checkPalindromeFormation(a string, b string) bool {
	n := len(a)
	if n == 1 {
		return true
	}

	var pre_a int = 0
	for pre_a = 0; pre_a <= n/2 && a[pre_a] == b[n-pre_a-1]; {
		pre_a++
	}

	if pre_a >= n/2 {
		return true
	}

	var pre_b int = 0
	for pre_b = 0; pre_b <= n/2 && b[pre_b] == a[n-pre_b-1]; {
		pre_b++
	}

	if pre_b >= n/2 {
		return true
	}

	var pre_i int = max(pre_a, pre_b)

	// fmt.Println(pre_i)
	var suf_i int = 0
	a_is_palindrome := true
	for suf_i = pre_i; suf_i <= n/2; {
		// fmt.Println(suf_i, n-suf_i-1)
		if a[suf_i] != a[n-suf_i-1] {
			a_is_palindrome = false
			break
		}
		suf_i++
	}

	b_is_palindrome := true
	for suf_i = pre_i; suf_i <= n/2; {
		// fmt.Println(suf_i, n-suf_i-1)
		if b[suf_i] != b[n-suf_i-1] {
			b_is_palindrome = false
			break
		}
		suf_i++
	}

	return a_is_palindrome || b_is_palindrome
}

func main() {
	var a string = "pvhmupgqeltozftlmfjjde"
	var b string = "yjgpzbezspnnpszebzmhvp"
	fmt.Println(checkPalindromeFormation(a, b))
}
