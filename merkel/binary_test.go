package merkel_test

import (
	"fmt"
	"testing"
	"github.com/Shubham-Rasal/merkel"
)

func TestBinary(t *testing.T) {
 
	root := merkel.NewNode(nil, nil, "1")
	root.Left = merkel.NewNode(nil, nil, "2")
	root.Right = merkel.NewNode(nil, nil, "3")
	root.Left.Left = merkel.NewNode(nil, nil, "4")
	root.Left.Right = merkel.NewNode(nil, nil, "5")
	root.Right.Left = merkel.NewNode(nil, nil, "6")


	merkel.Inorder(root)
	fmt.Println()
}