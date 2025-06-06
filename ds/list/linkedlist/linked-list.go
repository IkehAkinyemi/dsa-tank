/*
	Package linkedlist implements the data structures that supports the `list` ADT description. 
	The data structures in this package will implement the supported operational part of the List ADT. 
	The data structures within this package includes:
		- Singly Linked List
		- Circular Linked List
*/
package linkedlist

// Node represent each node in a linked list.
type Node[T comparable] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}