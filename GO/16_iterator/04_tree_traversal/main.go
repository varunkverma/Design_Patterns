package main

import "fmt"

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
	n := &Node{
		Value: value,
		left:  left,
		right: right,
	}
	left.parent = n
	right.parent = n
	return n
}

func NewLeafNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

// left root right
type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	iot := &InOrderIterator{
		root:          root,
		Current:       root,
		returnedStart: false,
	}

	// Since, in Inorder traversal we start from the left most node, we make the current set to the left more Node
	for iot.Current.left != nil {
		iot.Current = iot.Current.left
	}

	return iot
}

func (iot *InOrderIterator) Reset() {
	iot.Current = iot.root
	iot.returnedStart = false
}

func (iot *InOrderIterator) MoveNext() bool {

	if iot.Current == nil {
		return false
	}

	if !iot.returnedStart {
		iot.returnedStart = true
		return true // can use first element
	}

	if iot.Current.right != nil {
		// if a right node exista, set the current to that
		iot.Current = iot.Current.right

		// now we can't just use the current as it might have left node(s) available, so move to th leftest node of the current node and make it currect
		for iot.Current.left != nil {
			iot.Current = iot.Current.left
		}
		return true
	} else {
		// no more right nodes, traverse back to the parent as long as current is the right node of the parent
		p := iot.Current.parent
		for p != nil && iot.Current == p.right {
			iot.Current = p
			p = p.parent
		}
		iot.Current = p

		// if tree is not fully traversed
		return iot.Current != nil
	}
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{
		root: root,
	}
}

func (bt *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(bt.root)
}

func main() {
	root := NewNode(1,
		NewLeafNode(2),
		NewLeafNode(3))

	bt := NewBinaryTree(root)

	for iot := bt.InOrder(); iot.MoveNext(); {
		fmt.Printf("%d, ", iot.Current.Value)
	}
	fmt.Println("\b")
}
