import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	hashTable := make(map[string][]string)
	for _, elem := range strs {
		runes := []rune(elem)
		slices.Sort(runes)
		sortString := string(runes)
		hashTable[sortString] = append(hashTable[sortString], elem)
	}

	for _, v := range hashTable {
		result = append(result, v)
	}

	return result 
}
