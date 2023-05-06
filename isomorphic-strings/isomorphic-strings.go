package isomorphic

// Runtime 0 ms Beats 100%
// Memory 2.6 MB Beats 98.86%
// import "fmt"

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	n := make(map[byte]byte)

	for i := range s {
		if value, ok := m[s[i]]; ok {
			if value != t[i] {
				return false
			}
		} else {
			// Case 2 characters in `s` map to same character in `t`
			// Eg: "badc" vs "baba"
			if _, ok := n[t[i]]; ok {
				return false
			}
			m[s[i]] = t[i]
			n[t[i]] = s[i]
		}
	}
	return true
}
