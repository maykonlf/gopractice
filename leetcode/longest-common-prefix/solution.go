package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
    prefixLength := 1
    response := ""

    actualPrefix := prefix(strs[0], prefixLength)
    for {
        for i := 1; i < len(strs); i++ {
            if prefix(strs[i], prefixLength) != actualPrefix {
                return response
            }
        }
        response = actualPrefix
        prefixLength+=1
        actualPrefix = prefix(strs[0], prefixLength)
        if response == actualPrefix {
            break
        }
    }

    return response
}

func prefix(s string, size int) string {
    if len(s) == 0 || len(s) <= size - 1 {
        return s
    }

    return s[:size]
}
