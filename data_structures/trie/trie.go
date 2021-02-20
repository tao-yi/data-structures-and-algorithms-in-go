package trie

type trieNode struct {
	children    []*trieNode
	isEndOfWord bool
}

func NewTrieNode() *trieNode {
	return &trieNode{}
}

// space efficient
// efficient prefix queries
type Trie interface {
	Insert()
}
type trie struct {
	root *trieNode
}
