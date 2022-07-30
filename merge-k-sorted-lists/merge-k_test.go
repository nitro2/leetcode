package mergeksortedlists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func ListNodeToString(l *ListNode) string {
	var outString string
	for l != nil {
		outString += fmt.Sprintf("%v -> ", l.Val)
		l = l.Next
	}
	return outString
}

func createLinkListFromArray(arr []int) *linkedList {
	p := &linkedList{}
	for _, v := range arr {
		p.PushBack(&ListNode{Val: v})
	}
	return p
}

func createArrayLinkedList(l [][]int) []*ListNode {
	c := make([]*ListNode, len(l))
	for i, _ := range l {
		c[i] = createLinkListFromArray(l[i]).head
	}
	return c
}

func TestRegression(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input:    [][]int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := mergeKLists(createArrayLinkedList(test.input))
			want := createLinkListFromArray(test.expected).head
			assert.Equal(t, ListNodeToString(want), ListNodeToString(got))
		})
	}
}
