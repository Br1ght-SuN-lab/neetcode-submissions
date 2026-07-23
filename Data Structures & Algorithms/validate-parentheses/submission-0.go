func isValid(s string) bool {
	stack := make([]byte, 0)
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}  

	for i := range len(s) {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		//close bracket
		if len(stack) == 0 {
			return false //нечего закрывать
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1] //убрали top элемент

		if top != pairs[c] {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
