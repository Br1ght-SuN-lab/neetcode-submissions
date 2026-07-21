func trap(height []int) int {
	n := len(height)

	prefix := make([]int, n)
	suffix := make([]int, n)

	//заполнили префиксы и суффиксы
	for i := 0; i < n; i++ {
		if i == 0 {
			prefix[i] = -1 //выше ограничения
			continue
		} 
		prefix[i] = max(prefix[i-1], height[i-1])
	}

	for i := n-1; i >= 0; i-- {
		if i == n-1 {
			suffix[i] = -1;
			continue
		} 
		suffix[i] = max(suffix[i+1], height[i+1])
	}
	//решение задачи
	var result int = 0
	for i := range n {
		if i == 0 || i == n-1 {
			continue
		}
		value := min(suffix[i], prefix[i]) - height[i]
		if value < 0 {
			continue
		}
		result += value
	}

	return result 
}
