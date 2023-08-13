package merkel

import (
	"golang.org/x/crypto/sha3"
)

type MerkelTree struct {
	RootNode *MerkelNode
}

type MerkelNode struct {
	Left  *MerkelNode
	Right *MerkelNode
	Data  []byte
}

func NewMerkelNode(left, right *MerkelNode, data []byte) *MerkelNode {
	mNode := MerkelNode{}

	if left == nil && right == nil {
		hash := sha3.Sum256(data)
		mNode.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha3.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}

func NewMerkelTree(data [][]byte) *MerkelTree {
	var nodes []MerkelNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, d := range data {
		node := NewMerkelNode(nil, nil, d)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkelNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkelNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}

	mTree := MerkelTree{&nodes[0]}

	return &mTree
}