package p0006

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	row := make([]string, len(s))
	r := 0
	down := false

	for _, c := range s {
		row[r] += string(c)
		if r == 0 || r == numRows-1 {
			down = !down
		}
		if down {
			r++
		} else {
			r--
		}
	}

	return strings.Join(row, "")
}
