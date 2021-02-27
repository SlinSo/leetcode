package p0008

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\v' || b == '\f' || b == '\r'
}

func myAtoi(s string) int {
	res := 0
	sgn := 1
	if s == "" {
		return 0
	}

	for ; s != "" && isSpace(s[0]); s = s[1:] {

	}
	if s == "" {
		return 0
	}

	if s[0] == '-' {
		sgn = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for ; s != ""; s = s[1:] {
		d := int(s[0] - '0')
		if d > 9 {
			return sgn * res
		}
		res = res*10 + d
		if sgn*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sgn*res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return sgn * res
}

// cheating using strconv.Atoi, boilerplate is needed to mimic C atoi behaviour
func myAtoiCheat(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	minus := false
	plus := false

	end := -1
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] == '-' {
				minus = true
				continue
			}

			if s[i] == '+' {
				plus = true
				continue
			}
		}

		if s[i] < '0' || s[i] > '9' {
			end = i
			break
		}
	}
	if end > -1 {
		s = s[0:end]
	}
	if plus || minus {
		s = s[1:]
	}
	if s == "" || (minus && plus) {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		if minus {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if minus {
		i = -i
	}

	if i < math.MinInt32 {
		return math.MinInt32
	}

	if i > math.MaxInt32 {
		return math.MaxInt32
	}

	return i
}
