package p0003

func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	m := make(map[byte]int)
	for j, i := 0, 0; j < len(s); j++ {
		if n, ok := m[s[j]]; ok {
			i = max(n, i)
		}
		maxLen = max(maxLen, j-i+1)
		m[s[j]] = j + 1
	}
	return maxLen
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
