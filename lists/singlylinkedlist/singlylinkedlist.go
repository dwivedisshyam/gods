package singlylinkedlist

import "fmt"

// ---------------------------------- LinkedList ----------------------------------

type List[T comparable] struct {
	Head *Node[T]
}

func New[T comparable]() *List[T] {
	return &List[T]{
		Head: nil,
	}
}

func (ll *List[T]) SearchNode(data T) *Node[T] {
	if ll.Head == nil {
		return nil
	}

	var temp = ll.Head

	for temp != nil {
		if temp.Data == data {
			return temp
		}

		temp = temp.Next
	}
	return nil
}

func (ll *List[T]) DeleteNode(data T) {
	if ll.Head == nil {
		return
	}

	var temp = ll.Head

	if temp.Data == data && temp.Next != nil {
		ll.Head = ll.Head.Next
		return
	} else if temp.Data == data {
		temp = nil
		return
	}

	var prev *Node[T] = nil
	for temp != nil {
		if temp.Data == data && temp.Next != nil {
			temp.Data = temp.Next.Data
			temp.Next = temp.Next.Next
			return
		} else if temp.Data == data && temp.Next == nil {
			prev.Next = nil
			return
		}
		prev = temp
		temp = temp.Next
	}
}

func (ll *List[T]) AddNodes(data ...T) {
	for _, d := range data {
		ll.AddNode(d)
	}
}

func (ll *List[T]) AddNode(data T) {
	node := NewNode(data)

	if ll.Head == nil {
		ll.Head = node
		return
	}

	var temp = ll.Head

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = node
}

func (ll *List[T]) Print() {
	if ll.Head == nil {
		return
	}

	var temp = ll.Head

	for temp != nil {
		fmt.Printf("%v ", temp.Data)
		temp = temp.Next
	}
	fmt.Println()
}

// ---------------------------------- Node ----------------------------------

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Next: nil,
	}
}
