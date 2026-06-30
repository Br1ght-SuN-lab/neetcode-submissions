func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string, len(strs))
	for _, s := range strs {
		var count [26]byte
		for _, el := range s {
			count[el - 'a']++
		}
		groups[count] = append(groups[count], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result 
}
