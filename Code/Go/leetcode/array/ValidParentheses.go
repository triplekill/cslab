package array

func isValid(s string) bool {

	brackets := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := brackets[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if 0 == len(stack) {
				return false
			}
			ch := stack[len(stack)-1]
			if s[i] == brackets[ch] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if 0 < len(stack) {
		return false
	}
	return true
}
