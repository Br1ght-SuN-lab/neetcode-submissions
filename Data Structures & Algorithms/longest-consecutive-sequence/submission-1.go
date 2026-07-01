import (
	"slices"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)
	maxSeq := 1
	curSeq := 1

	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == 1 {
			curSeq++
		} else if nums[i+1] == nums[i] {
			continue
		} else {
			maxSeq = max(curSeq, maxSeq)
			curSeq = 1
		}
	}

	maxSeq = max(curSeq, maxSeq)

	return maxSeq
}
