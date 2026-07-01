func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))

	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 1
	suffix[len(nums)-1] = 1
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(suffix)-2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result = append(result, prefix[i] * suffix[i])
	}

	return result 
}
