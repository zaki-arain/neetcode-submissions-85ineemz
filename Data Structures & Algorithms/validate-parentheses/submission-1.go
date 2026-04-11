func isValid(str string) bool {
    stack := []rune{}
    brackets := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, s := range str {
        v, isClosing := brackets[s]; if isClosing {
            if len(stack) == 0 {return false}
            
            // closed bracket, check if matches top of open bracket stack
            top := stack[len(stack)-1]
            if top != v {return false}

            stack = stack[:len(stack)-1]
            continue
        }

        // opening bracket, push to stack
        stack = append(stack, s)
    }

    return len(stack) == 0
}
