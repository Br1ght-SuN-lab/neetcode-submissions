func maxArea(heights []int) int {
	n := len(heights)
	left := 0
	right := n - 1
	
	var result int = 0
	var prev int = -1
	var flag int = 0
	for left < right {
		minn := min(heights[left], heights[right])
		if prev != -1 && minn < prev {
			if flag == -1 {
				right--
			} else if flag == 1 {
				left++
			} 
			continue
		}
		curArea := minn * (right - left)
		result = max(result, curArea)

		if minn == heights[left] {
			prev = minn
			left++
			flag = 1
		} else {
			prev = minn
			right--
			flag = -1
		}

		fmt.Printf("left:%d, right:%d \n", left, right)
	}

	return result
}
