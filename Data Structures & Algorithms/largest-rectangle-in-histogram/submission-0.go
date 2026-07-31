func largestRectangleArea(heights []int) int {
	stack := make([][2]int, 0) //pairs: index, heights
	maxArea := 0
	for i, h := range heights {
		if len(stack) == 0 {
			stack = append(stack, [2]int{i, h}) 
			continue
		}

		popInd := -1
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if h >= top[1] {
				break
			}
			area := (i - top[0]) * top[1]
			maxArea = max(maxArea, area)
			popInd = top[0]
			stack = stack[:len(stack)-1]
		}

		var addInd int
		if popInd == -1 {
			addInd = i
		} else {
			addInd = popInd
		}
		stack = append(stack, [2]int{addInd, h})
	} 

	for _, elem := range stack {
		maxArea = max(maxArea, (len(heights) - elem[0]) * elem[1])
	}	

	return maxArea
}
