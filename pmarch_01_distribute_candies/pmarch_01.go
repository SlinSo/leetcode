package pmarch_01

func distributeCandies(candyType []int) int {
	upperBound := len(candyType) / 2

	m := make(map[int]struct{})
	for _, v := range candyType {
		m[v] = struct{}{}
	}

	numTypes := len(m)
	return min(numTypes, upperBound)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
