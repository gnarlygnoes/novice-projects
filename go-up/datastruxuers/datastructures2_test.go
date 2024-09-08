package parser

import (
	"fmt"
	"testing"
)

type DLList struct {
	Value any
	Prev  *DLList
	Next  *DLList
}

func (l *DLList) Append(value any) {
	if l.Next == nil {
		l.Next = &DLList{
			Value: value,
			Prev:  l,
			Next:  nil,
		}
	} else {
		l.Next.Append(value)
	}
}

func (l *DLList) Prepend(value any) {
	if l.Prev == nil {
		l.Prev = &DLList{
			Value: value,
			Prev:  nil,
			Next:  l,
		}
	} else {
		l.Prev.Prepend(value)
	}
}

func (l *DLList) GetFirst() *DLList {
	if l.Prev == nil {
		return l
	}
	return l.Prev.GetFirst()
}

func (l *DLList) Size() int {
	current := l.GetFirst()
	size := 1

	for current.Next != nil {
		current = current.Next
		size++
	}
	return size
}

func (l *DLList) PrintValues() {
	fmt.Print(l.Value, " ")
	if l.Next != nil {
		l.Next.PrintValues()
	}
}

func TestDLList(t *testing.T) {
	list := &DLList{
		Value: 0,
		Prev:  nil,
		Next:  nil,
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Prepend(-1)
	list.Prepend(-2)

	list.GetFirst().PrintValues()

	fmt.Println("Size: ", list.Size())
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func (t *Tree) Insert(value int) {
	if value < t.Value {
		if t.Left == nil {
			t.Left = &Tree{
				Value: value,
			}
		} else {
			t.Left.Insert(value)
		}
	} else if value > t.Value {
		if t.Right == nil {
			t.Right = &Tree{
				Value: value,
			}
		} else {
			t.Right.Insert(value)
		}
	}
}

func (t *Tree) Contains(value int) bool {
	if t.Value == value {
		return true
	}
	if value < t.Value {
		if t.Left == nil {
			return false
		}
		return t.Left.Contains(value)
	} else {
		if t.Right == nil {
			return false
		}
		return t.Right.Contains(value)
	}
}

func TestTree(t *testing.T) {
	nums := []int{5, 10, 1, 2, 5, 6, 3, 3, 3, 7, 8, 8, 9}

	tree := &Tree{
		Value: 5,
	}

	for _, n := range nums {
		tree.Insert(n)
	}

	fmt.Println(tree.Contains(10), tree.Contains(4))
}
