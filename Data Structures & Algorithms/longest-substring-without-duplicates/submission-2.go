func lengthOfLongestSubstring(s string) int {
	n := len(s)
	hashMap := make(map[byte]int)

	left := 0
	var result int = 0
	for right := 0; right < n; right++ {
		sym := s[right]

		if ind, ok := hashMap[sym]; ok && ind >= left {
			left = ind + 1
		}

		hashMap[sym] = right
		result = max(result, right - left + 1)
	}

	return result
}
