package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func NewList() List {
	return &list{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.len++
	item := &ListItem{
		Value: v,
		Next:  l.head,
	}
	if l.head != nil {
		l.head.Prev = item
	}
	l.head = item

	if l.tail == nil {
		l.tail = item
	}

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.len++
	item := &ListItem{
		Value: v,
		Prev:  l.tail,
	}
	if l.tail != nil {
		l.tail.Next = item
	}
	l.tail = item

	if l.head == nil {
		l.head = item
	}

	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	l.len--
	i = nil
}

func (l *list) MoveToFront(i *ListItem) {
	if l.head == i {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}
	i.Prev = nil
	i.Next = l.head
	l.head = i
}
