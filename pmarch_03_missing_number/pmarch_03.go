package pmarch_03

import "sort"

func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	s := 0
	for _, v := range nums {
		s += v
	}

	return sum - s
}

func missingNumber(nums []int) int {
	missing := len(nums)

	for i := 0; i < missing; i++ {
		missing ^= i ^ nums[i]
	}

	return missing
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
