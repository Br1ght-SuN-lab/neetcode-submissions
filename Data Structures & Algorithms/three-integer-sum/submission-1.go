import (
	"slices"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	slices.Sort(nums)

	n := len(nums)

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1
		if nums[i] > 0 || (i > 0 && nums[i-1] == nums[i]) {
			continue
		}

		target := 0 - nums[i]
		sum := nums[left] + nums[right]
		for left < right {
			sum = nums[left] + nums[right]
			if (sum < target) {
				left++
			} else if (sum > target) {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left]==nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return result
}
