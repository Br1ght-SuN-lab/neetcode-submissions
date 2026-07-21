func lengthOfLongestSubstring(s string) int {
	result := 0
	hashMap := make(map[rune]bool)
	runes := []rune(s)
	left := 0
	right := 0

	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	for left < len(s) && right < len(s) {
		if !hashMap[runes[right]] {
			hashMap[runes[right]] = true 
			right++
		} else {
			result = max(result, right - left)
			hashMap[runes[left]] = false 
			left++
		}

		fmt.Printf("left: %d, right: %d\n", left, right)
	}

	result = max(result, right - left)
	return result 
}
