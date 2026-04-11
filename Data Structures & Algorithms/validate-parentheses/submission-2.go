func isValid(str string) bool {
    stack := []rune{}
    brackets := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, s := range str {
        match, isClosing := brackets[s]; if isClosing {
            // closed bracket, check if matches top of open bracket stack
            if len(stack) == 0 || stack[len(stack)-1] != match {
                return false
            }
            
            // matched, so remove paired opening bracket from stack
            stack = stack[:len(stack)-1]
            continue
        }

        // opening bracket, push to stack
        stack = append(stack, s)
    }

    return len(stack) == 0
}
