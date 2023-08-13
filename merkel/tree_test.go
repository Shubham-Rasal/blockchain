package merkel_test

import (
	"fmt"
	"testing"

	"github.com/Shubham-Rasal/blockchain/merkel"
)


func TestMerkelTree(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world")}
	mTree := merkel.NewMerkelTree(data)
	fmt.Println(mTree.RootNode.Data)
	
}