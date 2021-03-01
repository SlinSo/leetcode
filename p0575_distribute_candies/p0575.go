package p0575

import "sort"

func distributeCandies(candyType []int) int {
	size := len(candyType)
	upperBound := size / 2
	numTypes := 1

	sort.Slice(candyType, func(i, j int) bool { return candyType[i] < candyType[j] })

	for i := 1; i < size; i++ {
		if candyType[i] != candyType[i-1] {
			numTypes++
		}
		if numTypes == upperBound {
			return numTypes
		}
	}

	return min(numTypes, upperBound)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
