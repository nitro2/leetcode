package long

func longestCommonPrefix(strs []string) string {
	nItems := len(strs)
	if nItems <= 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < nItems; j++ {
			if i < len(strs[j]) {
				// fmt.Printf("%c %c\n", strs[0][i], strs[j][i])
				if strs[0][i] != strs[j][i] {
					return strs[0][:i]
				}
			} else {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
