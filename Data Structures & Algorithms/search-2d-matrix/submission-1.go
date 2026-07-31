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
	arr := make([]int, len(matrix)*len(matrix[0]))
	fmt.Println(arr)
	for i, row := range matrix {
		for j, elem := range row {
			arr[i * len(row) + j] = elem
		} 
	}

	result := binarySearch(arr, target)
	return result
}
