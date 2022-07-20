package consecutive

import "fmt"

func consecutiveNumbersSum(n int) int {
	if n <= 2 {
		return 1
	}
	// x = n/r - (r-1)/2  --> Find integer of x
	count := 1
	x := 0.0
	for r := 2.0; r < float64(n) && x >= 0; r++ {
		x = (float64(n)/r - float64(r-1)/2)
		// fmt.Println(n, x, r)
		if x > 0 && x == float64(int(x)) {
			fmt.Println(n, x, r)
			count++
		}
	}
	return count
}
