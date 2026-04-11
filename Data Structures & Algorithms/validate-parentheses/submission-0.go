func isValid(s string) bool {
    if len(s)%2 != 0 {
        return false
    }
    
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    stack := []rune{}
    
    for _, r := range s {
    	if openBracket, isClosing := pairs[r]; isClosing {
            if len(stack) == 0 || stack[len(stack)-1] != openBracket {
                return false
            }
            // Pop from stack
            stack = stack[:len(stack)-1]
        } else {
            // It's an opening bracket
            stack = append(stack, r)
        }
    }
    
    return len(stack) == 0
}