package lists

import "unsafe"

type ListItem[T any] struct {
	next *ListItem[T]
	prev *ListItem[T]
	head *ListItem[T]
}

func (it *ListItem[T]) UnsafeGet() *T {
	return (*T)(unsafe.Pointer(it))
}

type ListHead[T any, P interface {
	*T
	getList() *ListItem[T]
}] struct {
	ListItem[T]
}

func (l *ListHead[T, P]) First() *T {
	return (*T)(unsafe.Pointer(l.next))
}

func (l *ListHead[T, P]) FirstSafe() *T {
	return l.NextSafe().UnsafeGet()
}

func (l *ListHead[T, P]) Clear() {
	l.next = &l.ListItem
	l.prev = &l.ListItem
	l.head = &l.ListItem
}

func (l *ListHead[T, P]) Append(p P) {
	if l == nil || p == nil || l.next == nil || l.prev == nil || l.head == nil {
		panic("nil list/item, or list uninitialized")
	}
	cur := p.getList()
	it := l.prev
	cur.next = &l.ListItem
	cur.prev = it
	l.prev = cur
	if it != nil { // see above note
		it.next = cur
	}
}

func (l *ListItem[T]) getList() *ListItem[T] {
	return l
}

func (l *ListItem[T]) Next() *ListItem[T] {
	return l.next
}

func (l *ListItem[T]) NextSafe() *ListItem[T] {
	if l == nil {
		return nil
	}
	it := l.next
	if it == it.head {
		return nil
	}
	return it
}

func (l *ListItem[T]) Remove() {
	l.prev.next = l.next
	l.next.prev = l.prev
	l.next = l
	l.prev = l
}
