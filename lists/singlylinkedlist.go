package lists

import (
	"fmt"
	"reflect"
)

// ---------------------------------- LinkedList ----------------------------------

type SinglyLinkedList[T any] struct {
	Head *SinglyNode[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Head: nil,
	}
}

func (ll *SinglyLinkedList[T]) SearchNode(data T) *SinglyNode[T] {
	if ll.Head == nil {
		return nil
	}

	var temp = ll.Head

	for temp != nil {
		if reflect.DeepEqual(temp.Data, data) {
			return temp
		}

		temp = temp.Next
	}
	return nil
}

func (ll *SinglyLinkedList[T]) DeleteNode(data T) {
	if ll.Head == nil {
		return
	}

	var temp = ll.Head

	if reflect.DeepEqual(temp.Data, data) && temp.Next != nil {
		ll.Head = ll.Head.Next
		return
	} else if reflect.DeepEqual(temp.Data, data) {
		temp = nil
		return
	}

	var prev *SinglyNode[T] = nil
	for temp != nil {
		if reflect.DeepEqual(temp.Data, data) && temp.Next != nil {
			temp.Data = temp.Next.Data
			temp.Next = temp.Next.Next
			return
		} else if reflect.DeepEqual(temp.Data, data) && temp.Next == nil {
			prev.Next = nil
			return
		}
		prev = temp
		temp = temp.Next
	}
}

func (ll *SinglyLinkedList[T]) AddNodes(data ...T) {
	for _, d := range data {
		ll.AddNode(d)
	}
}

func (ll *SinglyLinkedList[T]) AddNode(data T) {
	node := NewSinglyNode(data)

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

func (ll *SinglyLinkedList[T]) Print() {
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

type SinglyNode[T any] struct {
	Data T
	Next *SinglyNode[T]
}

func NewSinglyNode[T any](data T) *SinglyNode[T] {
	return &SinglyNode[T]{
		Data: data,
		Next: nil,
	}
}
