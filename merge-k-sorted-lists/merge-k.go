/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package mergeksortedlists

func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	} else if size == 1 {
		return lists[0]
	} else if size == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else {
		list1 := mergeKLists(lists[0 : size/2])
		list2 := mergeKLists(lists[size/2:])
		return mergeTwoLists(list1, list2)
	}
}

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
