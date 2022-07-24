// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.3 MB, less than 15.38% of Go online submissions for Valid Parentheses.
package validParentheses

// This experiment the builtin list of Golang
import (
	list "container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

func (l *Stack) Push(e any) {
	// fmt.Printf("Push %v \n", e)
	l.stack.PushFront(e)
}

func (l *Stack) Pop() any {

	if l.stack.Len() == 0 {
		return nil
	}
	frontElement := l.stack.Front()
	e := l.stack.Remove(frontElement)
	// fmt.Printf("Pop %v \n", e)
	return e
}

func isValid(s string) bool {
	l := &Stack{stack: list.New()}
	for _, c := range s {
		switch c {
		// Push element to stack
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '(':
			l.Push(string(c))
		// Pop element from stack then check if they are ok
		case ']':
			i := l.Pop()
			if fmt.Sprint(i) != "[" {
				return false
			}
		case '}':
			i := l.Pop()
			if fmt.Sprint(i) != "{" {
				return false
			}
		case ')':
			i := l.Pop()
			if fmt.Sprint(i) != "(" {
				return false
			}
		}
	}
	// fmt.Printf("Len of stack: %v", l.stack.Len())
	if l.stack.Len() > 0 {
		return false
	} else {
		return true
	}
}
