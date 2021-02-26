package p0005

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		l1 := score(s, i, i)
		l2 := score(s, i, i+1)
		l := max(l1, l2)
		fmt.Println(l1, l2, l)
		if l > end-start+1 {
			start = i - (l-1)/2
			end = i + l/2
		}
	}

	return s[start : end+1]
}

func max(i, j int) int {
	if j > i {
		return j
	}
	return i
}

func score(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
