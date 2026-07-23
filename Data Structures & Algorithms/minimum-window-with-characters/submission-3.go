func getCharInd(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}

	return int(c - 'A') + 26
}


func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := [52]int{}
	for i := 0; i < len(t); i++ {
		need[getCharInd(t[i])]++
	}
	required := 0
	for _, v := range need {
		if v > 0 {
			required++
		}
	}

	current := [52]int{}
	formed := 0 // сколько различных нужных символов уже набрано в достаточном количестве
	left := 0
	bestLen := -1
	bestStart := 0

	for right := 0; right < len(s); right++ {
		ind := getCharInd(s[right])
		current[ind]++
		if need[ind] > 0 && current[ind] == need[ind] {
			formed++
		}

		for formed == required {
			if bestLen == -1 || right-left+1 < bestLen {
				bestLen = right - left + 1
				bestStart = left
			}
			lind := getCharInd(s[left])
			current[lind]--
			if need[lind] > 0 && current[lind] < need[lind] {
				formed--
			}
			left++
		}
	}

	if bestLen == -1 {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}