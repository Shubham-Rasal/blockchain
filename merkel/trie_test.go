package merkel_test

import (
	"fmt"
	"testing"

	"github.com/Shubham-Rasal/merkel"
)


func TestInsertion(t *testing.T) {
	fmt.Println("Hello, World!")
	root := merkel.TrieNode{
		Children: [256]*merkel.TrieNode{},
		IsLeaf:   false,
	}
	merkel.InsertWord("he", &root)	

	if root.Children['h'] == nil {
		t.Error("Insertion failed")
	}

	if root.Children['h'].Children['e'] == nil {
		t.Error("Insertion failed")
	}

}

func TestSearch(t *testing.T) {
	fmt.Println("Hello, World!")
	root := merkel.TrieNode{
		Children: [256]*merkel.TrieNode{},
		IsLeaf:   false,
	}
	merkel.InsertWord("he", &root)	

	if !merkel.SearchWord("he", &root) {
		t.Error("Search failed")
	}

	if merkel.SearchWord("hello", &root) {
		t.Error("Search failed")
	}

}