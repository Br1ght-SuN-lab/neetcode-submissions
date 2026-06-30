type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result strings.Builder
	for _, str := range strs {
		result.WriteString(strconv.Itoa(len(str)))
		result.WriteByte('#')
		result.WriteString(str)
	}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		j := i 
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		result = append(result, encoded[j+1: j+length+1])
		i = j + length + 1
	}

	return result 
}
