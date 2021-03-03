package p0268

import "sort"

func missingNumber(nums []int) int {
	missing := len(nums)

	for i, n := range nums {
		missing ^= i ^ n
	}

	return missing
}

func missingNumberSum(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	s := 0
	for _, v := range nums {
		s += v
	}

	return sum - s
}

func missingNumberSort(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
