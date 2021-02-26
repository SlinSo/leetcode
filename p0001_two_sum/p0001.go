package p0001_two_sum

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
