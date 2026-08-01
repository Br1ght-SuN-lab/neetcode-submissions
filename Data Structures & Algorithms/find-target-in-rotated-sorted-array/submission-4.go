func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			if (target > nums[mid] || target <= nums[right]) {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] >= nums[left] {
				if nums[mid] < target {
					left = mid + 1
				} else if nums[mid] > target {
					right = mid 
				} else {
					return mid
				}
			} else {
				if target > nums[mid] && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid
				}
			}
		}
	}

	if nums[left] != target {
		return -1
	}
	return left
}
