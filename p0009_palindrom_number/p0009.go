package p0009

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return x == reverted || x == reverted/10
}

func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	b := strconv.Itoa(x)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}

	return true
}
