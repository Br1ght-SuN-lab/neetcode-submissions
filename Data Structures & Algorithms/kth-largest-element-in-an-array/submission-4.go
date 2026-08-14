func binSearch(data []int, val int) int {
	left := 0
	right := len(data)
	for left < right {
		mid := (left + right) / 2
		if data[mid] < val {
			left = mid + 1
		} else if data[mid] > val {
			right = mid
		} else {
			return mid
		}
 	}

	return left 
}


func insertValue(data *[]int, value int) {
	if len(*data) == 0 {
		*data = []int{value}
		return 
	}

	insertPos := binSearch(*data, value)
	*data = append(*data, 0)
	copy((*data)[insertPos+1:], (*data)[insertPos:])
	(*data)[insertPos] = value 
}

func findKthLargest(nums []int, k int) int {
	data := []int{}
	for len(data) < k {
		insertValue(&data, nums[0])
		nums = nums[1:]
	}

	for _, el := range nums {
		if el < data[0] {
			continue
		} 

		insertValue(&data, el)
		data = data[1:]
	}

	return data[0]
}
