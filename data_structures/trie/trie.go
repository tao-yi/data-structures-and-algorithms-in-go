package trie

// the number of possible characters in the trie
const AlphabetSize = 26

type node struct {
	children [26]*node
	isEnd    bool
}

// space efficient
// efficient prefix queries
type Trie interface {
	Insert(string)
	Search(string) bool
}

type trie struct {
	root *node
}

func NewTrie() Trie {
	return &trie{
		root: &node{},
	}
}

// insert a word into the trie
func (t *trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := int(w[i] - 'a')
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// search will take in a word and return true if that word is included inthe trie
func (t *trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := int(w[i] - 'a')
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	// if the last char is the end => exact match, return true
	// else => not an exact match, there return false
	return currentNode.isEnd
}
