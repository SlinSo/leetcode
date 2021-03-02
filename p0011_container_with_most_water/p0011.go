package p0011

func maxArea(height []int) int {
	l := len(height)
	max := area(height, 0, l-1)

	for i, j := 0, l-1; i < j; {
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}

		t := area(height, i, j)
		if t > max {
			max = t
		}
	}

	return max
}

func area(height []int, a, b int) int {
	h := min(height[a], height[b])

	return h * (b - a)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxAreaFirst(height []int) int {
	l := len(height)
	max := area(height, 0, l-1)

	hLeft := height[0]
	hRight := height[l-1]

	for i, j := 0, l-1; i < j; {
		if hLeft <= hRight {
			for height[i] <= hLeft && i < j {
				i++
			}
			hLeft = height[i]
		} else if hRight <= hLeft {
			for height[j] <= hRight && i < j {
				j--
			}
			hRight = height[j]
		}

		t := area(height, i, j)
		if t > max {
			max = t
		}
	}

	return max
}
