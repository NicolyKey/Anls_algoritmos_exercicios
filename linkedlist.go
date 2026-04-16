package main

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) InsertEnd(value int) {
	node := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}
	ll.Size++
}

func (ll *LinkedList) InsertAt(index, value int) {
	node := &Node{Value: value}
	if index == 0 {
		node.Next = ll.Head
		ll.Head = node
		if ll.Size == 0 {
			ll.Tail = node
		}
	} else {
		current := ll.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		node.Next = current.Next
		current.Next = node
		if node.Next == nil {
			ll.Tail = node
		}
	}
	ll.Size++
}

func (ll *LinkedList) RemoveFirst() {
	if ll.Head == nil {
		return
	}
	ll.Head = ll.Head.Next
	ll.Size--
	if ll.Size == 0 {
		ll.Tail = nil
	}
}

func (ll *LinkedList) RemoveLast() {
	if ll.Head == nil {
		return
	}
	if ll.Size == 1 {
		ll.Head = nil
		ll.Tail = nil
		ll.Size = 0
		return
	}
	current := ll.Head
	for current.Next != ll.Tail {
		current = current.Next
	}
	current.Next = nil
	ll.Tail = current
	ll.Size--
}

func (ll *LinkedList) RemoveAt(index int) {
	if index == 0 {
		ll.RemoveFirst()
		return
	}
	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	removed := current.Next
	current.Next = removed.Next
	if removed == ll.Tail {
		ll.Tail = current
	}
	ll.Size--
}

func (ll *LinkedList) Get(index int) int {
	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}
