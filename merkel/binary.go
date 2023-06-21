package merkel

import (
	"fmt"
)

type Node struct {
	Left *Node
	Right *Node
	Hash string
}

func NewNode(left *Node, right *Node, hash string) *Node {
	return &Node{left, right, hash}
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Print(root.Hash)
	Inorder(root.Right)
}

