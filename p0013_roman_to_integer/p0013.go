package p0013

import (
	"strings"
)

func romanToInt(s string) int {
	r := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	ans := 0

	for j := 0; j < len(r); j++ {
		for strings.HasPrefix(s, r[j]) {
			ans += n[j]
			s = s[len(r[j]):]
			if s == "" {
				return ans
			}
		}
	}

	return ans
}
