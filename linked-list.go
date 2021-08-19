package main

import (
	"fmt"
)

type MyLinkedList struct {
	head *Node
	n int
}

type Node struct {
	val int
	next *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index > this.n {
		return -1
	}

	var (
		iter = this.head
		idx int
	)

	for iter != nil {
		if idx == index {
			return iter.val
		}

		idx++
		iter = iter.next
	}

	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	node := &Node{
		val: val,
		next: this.head,
	}

	this.head = node
	this.n++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	node := &Node{
		val: val,
	}

	iter := this.head

	if this.n == 0 {
		this.head = &Node{
			val: val,
		}
		this.n++

		return
	}

	for iter.next != nil {
		iter = iter.next
	}

	iter.next = node
	this.n++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.Len() {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Len() {
		this.AddAtTail(val)
		return
	}

	node := &Node{
		val: val,
	}

	var (
		iter = this.head
		idx int
	)

	for iter != nil {
		if idx == index-1 {
			node.next = iter.next
			iter.next = node
			this.n++

			return
		}

		idx++
		iter = iter.next
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index > this.n-1 {
		return
	}

	var (
		iter = this.head
		prev *Node
		idx int
	)

	for iter != nil {
		if idx == index {
			if prev != nil {
				prev.next = iter.next
				this.n--
			} else {
				this.head = iter.next
			}

			return
		}

		idx++
		prev = iter
		iter = iter.next
	}
}

func (this *MyLinkedList) Len() int {
	return this.n
}

func (this *MyLinkedList) Print() {

	iter := this.head

	for iter != nil {
		fmt.Print(iter.val, " ")
		iter = iter.next
	}

	fmt.Println("len", this.Len())
}