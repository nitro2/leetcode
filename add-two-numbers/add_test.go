package add

import (
	"testing"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head *ListNode
}

func (l *linkedList) PushBack(n *ListNode) {
	if l.head == nil {
		l.head = n
	} else {
		h := l.head
		t := h.Next
		for t != nil {
			h = h.Next
			t = h.Next
		}
		h.Next = n
	}
}

func DisplayNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%v -> ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func TestAdd(t *testing.T) {
	el := &linkedList{}
	el.PushBack(&ListNode{Val: 4})
	el.PushBack(&ListNode{Val: 2})
	DisplayNode(el.head)

	in1 := &linkedList{}
	in1.PushBack(&ListNode{Val: 1})
	in1.PushBack(&ListNode{Val: 1})
	DisplayNode(in1.head)

	in2 := &linkedList{}
	in2.PushBack(&ListNode{Val: 3})
	in2.PushBack(&ListNode{Val: 1})
	DisplayNode(in2.head)

	out := addTwoNumbers(in1.head, in2.head)
	DisplayNode(out)

}
