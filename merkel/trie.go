package merkel

type TrieNode struct {
	Children [256]*TrieNode
	IsLeaf   bool
}

func InsertWord(word string , root *TrieNode) {
	
  	for i := 0; i < len(word); i++ {
		if root.Children[word[i]] == nil {
			root.Children[word[i]] = new(TrieNode)
		}
		root = root.Children[word[i]]
	}

	root.IsLeaf = true
}


func SearchWord(word string , root *TrieNode) bool {
	
	for i := 0; i < len(word); i++ {
		if root.Children[word[i]] == nil {
			return false
		}
		root = root.Children[word[i]]
	}

	return root.IsLeaf
}
