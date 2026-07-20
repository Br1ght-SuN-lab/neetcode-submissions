func maxArea(heights []int) int {
	var result int

	left, right := 0, len(heights)-1
	for left < right {
		r := (right - left) * min(heights[left], heights[right])
		if r > result {
			result = r
		}
		if heights[left] > heights[right] {
			right -= 1
		} else {
			left += 1
		}
	}

	return result
}
