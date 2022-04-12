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
	length    int
	frontItem *ListItem
	backItem  *ListItem
}

func NewList() List {
	return &list{}
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.frontItem
}

func (l list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.length == 0 {
		l.frontItem = newItem
		l.backItem = newItem
	} else {
		l.frontItem.Prev = newItem
		newItem.Next = l.frontItem
		l.frontItem = newItem
	}

	l.length++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.length == 0 {
		l.frontItem = newItem
		l.backItem = newItem
	} else {
		l.backItem.Next = newItem
		newItem.Prev = l.backItem
		l.backItem = newItem
	}

	l.length++

	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.frontItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.backItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	if i.Next == nil {
		l.backItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.frontItem.Prev = i
	i.Next = l.frontItem
	i.Prev = nil

	l.frontItem = i
}
