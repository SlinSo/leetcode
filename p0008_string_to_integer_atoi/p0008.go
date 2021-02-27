package p0008

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// cheating using strconv.Atoi, boilerplate is needed to mimic C atoi behaviour
func myAtoi(s string) int {
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
