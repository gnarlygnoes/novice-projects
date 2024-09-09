package parser

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func TestMap(t *testing.T) {
	m := map[string]Employee{}

	e := Employee{
		Name:   "Nigel",
		Age:    39,
		Salary: 45_000,
	}

	e2 := Employee{
		Name:   "Harry",
		Age:    34,
		Salary: 90_000,
	}

	e3 := Employee{
		Name:   "Jill",
		Age:    34,
		Salary: 90_000,
	}

	m[e.Name] = e
	m[e2.Name] = e2
	m[e3.Name] = e3

	jill, found := m["Jill"]

	if found {
		fmt.Println(jill)
	}

	gary, found := m["Gary"]

	if found {
		fmt.Println(gary)
	} else {
		fmt.Println("gary not found")
		fmt.Println(gary)
	}

	delete(m, "Larry")
	delete(m, "Harry")

	for name, employee := range m {
		fmt.Println(name, employee)
	}
}

type List struct {
	Value any
	Next  *List
}

func (l *List) Prepend(value any) *List {
	elem := List{
		Value: value,
		Next:  l,
	}

	return &elem
}

func (l *List) Append(value any) {
	if l.Next == nil {
		l.Next = &List{Value: value}
	} else {
		l.Next.Append(value)
	}
}

func (l *List) PrintValues() {
	fmt.Print(l.Value, " ")
	if l.Next != nil {
		l.Next.PrintValues()
	}
}

func TestList(t *testing.T) {
	list := &List{
		Value: 0,
		Next:  nil,
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list = list.Prepend(-1)
	list = list.Prepend("PEN15")

	list.PrintValues()
}

type List2[T any] struct {
	Value T
	Next  *List2[T]
}

func (l *List2[T]) Prepend(value T) *List2[T] {
	elem := List2[T]{
		Value: value,
		Next:  l,
	}

	return &elem
}

func (l *List2[T]) Append(value T) {
	if l.Next == nil {
		l.Next = &List2[T]{Value: value}
	} else {
		l.Next.Append(value)
	}
}

func (l *List2[T]) PrintValues() {
	fmt.Print(l.Value, " ")
	if l.Next != nil {
		l.Next.PrintValues()
	}
}

func TestList2(t *testing.T) {
	list := &List2[int]{
		Value: 0,
		Next:  nil,
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list = list.Prepend(-1)

	list.PrintValues()
}
