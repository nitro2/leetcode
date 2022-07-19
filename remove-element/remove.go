// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
// Memory Usage: 2 MB, less than 82.93% of Go online submissions for Remove Element
package remove

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
			i--
		}
	}
	return n
}
