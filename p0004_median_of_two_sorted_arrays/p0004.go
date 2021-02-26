package p0004

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ans := 0.0

	s := append(nums1, nums2...)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

	if len(s)%2 == 0 {
		nUpper := int(math.Ceil(float64(len(s)) / 2.0))
		ans = float64(s[nUpper]+s[nUpper-1]) / 2.0
	} else {
		middle := int(math.Floor(float64(len(s)) / 2.0))
		ans = float64(s[middle])
	}

	return ans
}
