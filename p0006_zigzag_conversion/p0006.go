package p0006

// with grid output
func convert(s string, numRows int) string {
	l := len(s)

	if numRows < 2 {
		return s
	}
	if l <= numRows {
		return s
	}

	cols := (l / numRows)
	cols += (cols*numRows - 2 + numRows)

	grid := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		grid[i] = make([]string, cols)
	}

	block := 2*numRows - 1
	nBlock := 0
	c := 0
	for i := 0; i < (nBlock+1)*block && i < l; i += (block - 1) {
		n := 0
		r := 0
		for r = 0; r < numRows && (i+n) < l; r++ {
			grid[r][c] = string(s[i+n])
			n++
		}
		r--
		for r > 1 && (i+n) < l {
			c++
			r--
			grid[r][c] = string(s[i+n])
			n++
		}
		c++
		nBlock++
	}

	ans := ""
	for i := 0; i < numRows; i++ {
		for c := 0; c < cols; c++ {
			ans += grid[i][c]
		}
	}

	return ans
}
