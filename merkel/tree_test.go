package merkel_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/Shubham-Rasal/blockchain/merkel"
)


func TestMerkelTree(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world")}
	mTree := merkel.NewMerkelTree(data)
	hash1 := sha256.Sum256([]byte("hello"))
	hash2 := sha256.Sum256([]byte("world"))
	expectedRoot := append(hash1[:], hash2[:]...)
	expectedRootHash := sha256.Sum256(expectedRoot)



	if check(mTree.RootNode.Data[:], expectedRootHash[:]) {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
	
}



func check(a, b []byte) bool {
	
	first := make([]int8, sha256.Size)

	for i := range a {
		first[i] = int8(a[i])
	}

	second := make([]int8, sha256.Size)

	for i := range b {
		second[i] = int8(b[i])
	}

	//check if the two slices are equal
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}	

	return true
}