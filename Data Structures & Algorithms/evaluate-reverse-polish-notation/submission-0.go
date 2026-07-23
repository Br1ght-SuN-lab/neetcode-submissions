func evalRPN(tokens []string) int {
	stack := []int{}

	for _, elem := range tokens {
		digit, ok := strconv.Atoi(elem)
		if ok != nil {
			second := stack[len(stack)-1]
			first := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			if elem == "+" {
				res = first + second
			} else if elem == "-" {
				res = first - second
			} else if elem == "*" {
				res = first * second
			} else if elem == "/" {
				res = first / second
			}

			stack = append(stack, res)
 		} else {
			stack = append(stack, digit)
		}
	}

	return stack[0]
}
