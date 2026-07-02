func binarySearch(numbers []int, start int, target int) int {
	left := start
	right := len(numbers)
	for left < right {
		middle := left + (right - left) / 2
		if numbers[middle] < target {
			left = middle + 1
		} else if numbers[middle] > target {
			right = middle
		} else {
			return middle
		}
	}

	return -1
}


func twoSum(numbers []int, target int) []int {
	for i, _ := range numbers {
		result := binarySearch(numbers, i+1, target-numbers[i])
		if result == -1 {
			continue
		}

		return []int{i+1, result+1}
	}

	return []int{}
}
