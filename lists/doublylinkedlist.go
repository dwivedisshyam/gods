package lists

import (
	"fmt"
	"reflect"
)

// ---------------------------- List ----------------------------
type DoublyLinkedList[T any] struct {
	Head *DoublyNode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Head: nil,
	}
}

func (ll *DoublyLinkedList[T]) InsertFirst(data T) {
	node := NewDoublyNode(data)
	ll.Head.Previous = node
	node.Next = ll.Head
	ll.Head = node
}

func (ll *DoublyLinkedList[T]) InsertLast(data T) {
	node := NewDoublyNode(data)
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

func (ll *DoublyLinkedList[T]) Delete(data T) {
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

func (ll *DoublyLinkedList[T]) DeleteFirst() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
	ll.Head.Previous = nil
}

func (ll *DoublyLinkedList[T]) DeleteLast() {
	if ll.Head == nil {
		return
	}

	temp := ll.Head
	for temp.Next != nil && temp.Next.Next != nil {
		temp = temp.Next
	}

	temp.Next = nil
}

func (ll *DoublyLinkedList[T]) PrintForward() {
	if ll.Head == nil {
		return
	}
	temp := ll.Head
	for temp != nil {
		fmt.Printf("%v ", temp.Data)
		temp = temp.Next
	}
}

func (ll *DoublyLinkedList[T]) PrintBackword() {
	if ll.Head == nil {
		return
	}
	ll.printBackword(ll.Head)
}

func (ll *DoublyLinkedList[T]) printBackword(node *DoublyNode[T]) {
	if node == nil {
		return
	}

	defer fmt.Printf("%v ", node.Data)

	ll.printBackword(node.Next)
}

// ---------------------------- Node ----------------------------
type DoublyNode[T any] struct {
	Data     T
	Next     *DoublyNode[T]
	Previous *DoublyNode[T]
}

func NewDoublyNode[T any](data T) *DoublyNode[T] {
	return &DoublyNode[T]{
		Data: data,
	}
}
