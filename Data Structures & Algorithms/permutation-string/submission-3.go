func checkInclusion(s1 string, s2 string) bool {
	dataFirst := [26]int{}
	for i := range len(s1) {
		dataFirst[s1[i] - 'a']++
	}

	left := 0 
	for r := range len(s2) {
		num := s2[r]-'a'
		dataFirst[num]--
		for dataFirst[num] < 0 {
			dataFirst[s2[left]-'a']++
			left++

		}

		if dataFirst == [26]int{} {
			return true 
		} 
	}
	return false 
}
