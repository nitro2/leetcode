/**
 * https://leetcode.com/problems/add-two-numbers/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package add

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := 0
	n := &ListNode{}
	var p *ListNode

	for l1 != nil && l2 != nil {
		if p == nil {
			p = n
		} else {
			t := &ListNode{}
			p.Next = t
			p = t
		}
		s := l1.Val + l2.Val + r
		if s > 10 {
			s = s % 10
			r = 1
		}
		p.Val = s

		l1 = l1.Next
		l2 = l2.Next
	}
	return n
}
