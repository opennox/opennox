package opennox

import "unsafe"

type listHead[T any, P interface {
	*T
	getList() *listItem
}] struct {
	listItem
}

func (l *listHead[T, P]) First() *T {
	return (*T)(unsafe.Pointer(l.Next()))
}

func (l *listHead[T, P]) Clear() {
	l.next = &l.listItem
	l.prev = &l.listItem
	l.head = &l.listItem
}

func (l *listHead[T, P]) Append(p P) {
	if l == nil || p == nil {
		panic("nil list or item")
	}
	cur := p.getList()
	it := l.prev

	// FIXME: in some cases 'it' is null, which suggests that some lists are not initialized properly
	//        auto-initializing them however leads to more serious issues like double-free
	if false {
		if it == nil && l.next == nil && l.head == nil {
			l.Clear()
			it = l.prev
		}
	}

	cur.next = &l.listItem
	cur.prev = it

	l.prev = cur
	if it != nil { // see above note
		it.next = cur
	}
}

type listItem struct {
	next *listItem
	prev *listItem
	head *listItem
}

func (l *listItem) getList() *listItem {
	return l
}

func (l *listItem) Next() *listItem {
	if l == nil {
		return nil
	}
	it := l.next
	if it == it.head {
		return nil
	}
	return it
}

func (l *listItem) Remove() {
	l.prev.next = l.next
	l.next.prev = l.prev
	l.next = l
	l.prev = l
}

func nox_common_list_getFirstSafe_425890(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer((*listHead[listItem, *listItem])(list).First())
}

func nox_common_list_getNextSafe_4258A0(list unsafe.Pointer) unsafe.Pointer {
	if list == nil {
		return nil
	}
	return nox_common_list_getNext_425940(list)
}

func nox_common_list_getNext_425940(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer((*listItem)(list).Next())
}

func nox_common_list_clear_425760(list unsafe.Pointer) {
	(*listHead[listItem, *listItem])(list).Clear()
}

func nox_common_list_append_4258E0(list, cur unsafe.Pointer) {
	(*listHead[listItem, *listItem])(list).Append((*listItem)(cur))
}
