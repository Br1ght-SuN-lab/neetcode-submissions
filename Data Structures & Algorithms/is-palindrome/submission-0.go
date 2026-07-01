import (
	"slices"
)

func isPalindrome(s string) bool {
	sentence := []string{}

	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			sentence = append(sentence, strings.ToLower(string(r)))
		}
	}

	fmt.Println(sentence)

	reverseSentence := slices.Clone(sentence)
	slices.Reverse(reverseSentence)
	if slices.Equal(sentence, reverseSentence) {
		return true
	}

	return false 
}
