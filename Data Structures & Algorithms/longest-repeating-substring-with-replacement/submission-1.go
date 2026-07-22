func characterReplacement(s string, k int) int {
	data := make(map[byte]int)
	left := 0

	n := len(s)
	maxf := 0
	result := 0
	for right := range n {
		data[s[right]]++
		maxf = max(maxf, data[s[right]])

		for (right - left + 1) - maxf > k {
			data[s[left]]--
			left++
		}

		result = max(right - left + 1)
	}

	return result 
}
