package bst

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}
