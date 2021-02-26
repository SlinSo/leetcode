package p0001

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
