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
	carry := 0
	n := &ListNode{}
	var p *ListNode

	for l1 != nil || l2 != nil || carry != 0 {
		var in1, in2 int
		if l1 != nil {
			in1 = l1.Val
			l1 = l1.Next
		} else {
			in1 = 0
		}
		if l2 != nil {
			in2 = l2.Val
			l2 = l2.Next
		} else {
			in2 = 0
		}

		if p == nil { // First entry we have to reuse the first node
			p = n
		} else { // Next entry we create new one
			t := &ListNode{Val: 0}
			p.Next = t
			p = t
		}
		s := in1 + in2 + carry
		if s >= 10 {
			s = s % 10
			carry = 1
		} else {
			carry = 0
		}
		p.Val = s

	}
	return n
}
