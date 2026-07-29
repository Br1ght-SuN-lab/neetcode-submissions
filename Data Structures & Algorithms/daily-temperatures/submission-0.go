func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	indexes := []int{}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if len(stack) <= 0 {
			stack = append(stack, temperatures[i])
			indexes = append(indexes, i)
			continue
		}

		top := stack[len(stack)-1]
		top_ind := indexes[len(indexes)-1]
		for temperatures[i] > top {
			result[top_ind] = i - top_ind
			stack = stack[:len(stack)-1]
			indexes = indexes[:len(indexes)-1]

			if len(stack) <= 0 {
				break
			}
			top = stack[len(stack)-1]
			top_ind = indexes[len(indexes)-1]
		}

		stack = append(stack, temperatures[i]) 
		indexes = append(indexes, i)
	}

	for _, ind := range indexes {
		result[ind] = 0
	}

	return result 
}
