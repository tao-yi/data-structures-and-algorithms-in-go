package trie_test

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/data_structures/trie"
)

func TestTrie(tc *testing.T) {
	t := trie.NewTrie()
	t.Insert("hello")
	fmt.Println(t.Search("he"))
	fmt.Println(t.Search("ab"))
	fmt.Println(t.Search("hell"))
	fmt.Println(t.Search("hello"))
	fmt.Println(t.Search("helloa"))
}
