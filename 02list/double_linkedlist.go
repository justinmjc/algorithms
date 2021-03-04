package main

type Element struct {
	prev, next *Element
	list       *List
	Value      interface{}
}

type List struct {
	root Element
	len  int
}

func New() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}
