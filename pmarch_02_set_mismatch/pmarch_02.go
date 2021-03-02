package pmarch_02

func findErrorNums(nums []int) []int {
	m := make(map[int]struct{})

	dup := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			dup = nums[i]
		}
		m[nums[i]] = struct{}{}
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, ok := m[i]; !ok {
			return []int{dup, i}
		}
	}

	return []int{}
}
