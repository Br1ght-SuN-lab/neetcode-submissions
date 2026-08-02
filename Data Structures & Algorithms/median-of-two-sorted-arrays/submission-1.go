func searchInsertionPosition(data []int, target int) int {
	left := 0
	right := len(data)
	for left < right {
		mid := (left + right) / 2
		if data[mid] < target {
			left = mid + 1
		} else if data[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
} 



func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, n := range nums2 {
		if len(nums1) == 0 {
			nums1 = append(nums1, n)
			continue
		}
		
		pos := searchInsertionPosition(nums1, n)
		if pos > len(nums1) - 1 && nums1[len(nums1) - 1] < n { 
			nums1 = append(nums1, n)
			continue
		}

		//вставка элемента из num2 в отсортрованном порядке в num1
		result := make([]int, 0, len(nums1)+1)
		result = append(result, nums1[:pos]...)
		result = append(result, n)
		result = append(result, nums1[pos:]...)
		nums1 = result
	}

	if len(nums1) % 2 == 0 {
		return (float64(nums1[len(nums1) / 2]) + float64(nums1[len(nums1) / 2 - 1])) / 2.0
	}

	return float64(nums1[len(nums1) / 2])
}
 