/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
Runtime: 0 ms, faster than 100.00% of Go online submissions
*/

package mergetwosortedlists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	out := &ListNode{}
	var h, p1, p2 *ListNode
	h = out
	p1 = list1
	p2 = list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			h.Next = &ListNode{Val: p1.Val}
			p1 = p1.Next
			h = h.Next
		} else {
			h.Next = &ListNode{Val: p2.Val}
			p2 = p2.Next
			h = h.Next
		}
	}
	for p1 != nil {
		h.Next = &ListNode{Val: p1.Val}
		p1 = p1.Next
		h = h.Next
	}
	for p2 != nil {
		h.Next = &ListNode{Val: p2.Val}
		p2 = p2.Next
		h = h.Next
	}
	return out.Next
}
