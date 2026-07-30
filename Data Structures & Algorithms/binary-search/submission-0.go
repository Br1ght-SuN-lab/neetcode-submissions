func search(nums []int, target int) int {
	left := 0 
	right := len(nums)
	for left < right {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}

	return -1
}
