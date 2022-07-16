package add

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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

func ListNodeToString(l *ListNode) string {
	var outString string
	for l != nil {
		outString += fmt.Sprintf("%v -> ", l.Val)
		l = l.Next
	}
	return outString
}

func TestAdd(t *testing.T) {
	el := &linkedList{}
	el.PushBack(&ListNode{Val: 4})
	el.PushBack(&ListNode{Val: 2})
	// DisplayNode(el.head)

	in1 := &linkedList{}
	in1.PushBack(&ListNode{Val: 1})
	in1.PushBack(&ListNode{Val: 1})
	// DisplayNode(in1.head)

	in2 := &linkedList{}
	in2.PushBack(&ListNode{Val: 3})
	in2.PushBack(&ListNode{Val: 1})
	// DisplayNode(in2.head)

	out := addTwoNumbers(in1.head, in2.head)

	// DisplayNode(out)
	assert.Equal(t, ListNodeToString(el.head), ListNodeToString(out))
}

func TestAddUnalign(t *testing.T) {
	in1 := &linkedList{}
	in1.PushBack(&ListNode{Val: 1})
	in1.PushBack(&ListNode{Val: 1})
	in1.PushBack(&ListNode{Val: 2})
	// DisplayNode(in1.head)

	in2 := &linkedList{}
	in2.PushBack(&ListNode{Val: 3})
	in2.PushBack(&ListNode{Val: 1})
	// DisplayNode(in2.head)

	el := &linkedList{}
	el.PushBack(&ListNode{Val: 4})
	el.PushBack(&ListNode{Val: 2})
	el.PushBack(&ListNode{Val: 2})
	// DisplayNode(el.head)

	out := addTwoNumbers(in1.head, in2.head)

	// DisplayNode(out)
	assert.Equal(t, ListNodeToString(el.head), ListNodeToString(out))
}

func createLinkListFromArray(arr []int) *linkedList {
	p := &linkedList{}
	for _, v := range arr {
		p.PushBack(&ListNode{Val: v})
	}
	return p
}

func TestRegression(t *testing.T) {
	tests := []struct {
		name     string
		in1      []int
		in2      []int
		expected []int
	}{
		{
			name:     "Input1 less element than input2",
			in1:      []int{1, 2},
			in2:      []int{2, 3, 4},
			expected: []int{3, 5, 4},
		},
		{
			name:     "Input1 more element than input2",
			in1:      []int{1, 2, 5},
			in2:      []int{2, 3},
			expected: []int{3, 5, 5},
		},
		{
			name:     "Carry result ",
			in1:      []int{5, 9},
			in2:      []int{6, 9, 4, 1},
			expected: []int{1, 9, 5, 1},
		},
		{
			name:     "Add 0",
			in1:      []int{0},
			in2:      []int{2, 3, 4},
			expected: []int{2, 3, 4},
		},
		{
			name:     "Add two 0",
			in1:      []int{0},
			in2:      []int{0},
			expected: []int{0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := createLinkListFromArray(test.in1)
			b := createLinkListFromArray(test.in2)
			exp := createLinkListFromArray(test.expected)

			out := addTwoNumbers(a.head, b.head)

			assert.Equal(t, ListNodeToString(exp.head), ListNodeToString(out))
		})
	}
}
