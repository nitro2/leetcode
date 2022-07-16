// Faster than 100.00% Go online submissions
package roman

func RomanToArabic(roman string) (res int) {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(roman) - 1; i >= 0; i-- {
		if dict[roman[i]] >= last {
			res += dict[roman[i]]
			last = dict[roman[i]]
		} else {
			res -= dict[roman[i]]
		}
	}

	return res
}
