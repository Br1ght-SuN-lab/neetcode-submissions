func longestConsecutive(nums []int) int {
	table := make(map[int]struct{})
	longest := 0

	for _, elem := range nums{
		table[elem] = struct{}{}
	}

	for _, num := range nums {
		if _, exist := table[num-1]; exist {
			continue
		}

		length := 1
		
		for {
			if _, exist := table[num+length]; !exist {
				break
			}
			length++
		}

		longest = max(longest, length)
	}

	return longest
}
