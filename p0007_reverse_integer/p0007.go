package p0007

import (
	"math"
	"strconv"
)

// int is int64
func reverse(x int) int {
	b := []byte(strconv.Itoa(x))
	minus := false

	if b[0] == '-' {
		b = b[1:]
		minus = true
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0
	}

	if minus {
		i = -i
	}

	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0
	}

	return i
}
