func binarySearch(array []int, target int) bool {
	left := 0 
	right := len(array)
	for left < right {
		mid := (left + right) / 2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid
		} else {
			return true 
		}
	}

	return false 
}


func searchMatrix(matrix [][]int, target int) bool {
	var result bool 
	for _, row := range matrix {
		result = binarySearch(row, target)
		if result == true {
			break
		}
	}

	return result
}
