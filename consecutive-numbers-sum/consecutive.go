// Runtime: 4 ms, faster than 100.00% of Go online submissions for Consecutive Numbers Sum.
// Memory Usage: 2 MB, less than 71.43% of Go online submissions for Consecutive Numbers Sum.
package consecutive

func consecutiveNumbersSum(n int) int {
	if n <= 2 {
		return 1
	}
	// x = n/r - (r-1)/2  --> Find integer of x
	count := 1
	x := 0.0
	for r := 2.0; r*r < float64(2*n) && x >= 0; r++ {
		x = (float64(n)/r - float64(r-1)/2)
		// fmt.Println(n, x, r)
		if x > 0 && x == float64(int(x)) {
			// fmt.Println(n, x, r)
			count++
		}
	}
	return count
}
