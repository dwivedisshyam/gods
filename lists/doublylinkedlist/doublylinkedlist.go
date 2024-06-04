package doublylinkedlist

import (
	"fmt"
	"reflect"
)

// ---------------------------- List ----------------------------
type List[T any] struct {
	Head *Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		Head: nil,
	}
}

func (ll *List[T]) InsertFirst(data T) {
	node := NewNode(data)
	ll.Head.Previous = node
	node.Next = ll.Head
	ll.Head = node
}

func (ll *List[T]) InsertLast(data T) {
	node := NewNode(data)
	if ll.Head == nil {
		ll.Head = node
		return
	}

	temp := ll.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
	node.Previous = temp
}

func (ll *List[T]) Delete(data T) {
	if ll.Head == nil {
		return
	}

	if reflect.DeepEqual(ll.Head.Data, data) {
		ll.DeleteFirst()
		return
	}

	temp := ll.Head
	for temp.Next != nil {
		if reflect.DeepEqual(temp.Data, data) {
			temp.Previous.Next = temp.Next
			temp.Next.Previous = temp.Previous
			return
		}
		temp = temp.Next
	}

	if reflect.DeepEqual(temp.Data, data) {
		ll.DeleteLast()
	}
}

func (ll *List[T]) DeleteFirst() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
	ll.Head.Previous = nil
}

func (ll *List[T]) DeleteLast() {
	if ll.Head == nil {
		return
	}

	temp := ll.Head
	for temp.Next != nil && temp.Next.Next != nil {
		temp = temp.Next
	}

	temp.Next = nil
}

func (ll *List[T]) PrintForward() {
	if ll.Head == nil {
		return
	}
	temp := ll.Head
	for temp != nil {
		fmt.Printf("%v ", temp.Data)
		temp = temp.Next
	}
}

func (ll *List[T]) PrintBackword() {
	if ll.Head == nil {
		return
	}
	ll.printBackword(ll.Head)
}

func (ll *List[T]) printBackword(node *Node[T]) {
	if node == nil {
		return
	}

	defer fmt.Printf("%v ", node.Data)

	ll.printBackword(node.Next)
}

// ---------------------------- Node ----------------------------
type Node[T any] struct {
	Data     T
	Next     *Node[T]
	Previous *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data: data,
	}
}
