package ds_test

import (
	"go-data-structure/ds"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() (bool, error)
		expectedRes bool
		expectedErr error
	}{
		{
			name: "test search word",
			fn: func() (bool, error) {
				rootNode := ds.TrieNode{
					Children: map[rune]*ds.TrieNode{},
				}
				trie := ds.Trie{Root: &rootNode}
				trie.Insert("test")
				found := trie.Search("test")
				return found, nil
			},
			expectedRes: true,
			expectedErr: nil,
		},
		{
			name: "test search word not found",
			fn: func() (bool, error) {
				rootNode := ds.TrieNode{
					Children: map[rune]*ds.TrieNode{},
				}
				trie := ds.Trie{Root: &rootNode}
				trie.Insert("test")
				found := trie.Search("testeez")
				return found, nil
			},
			expectedRes: false,
			expectedErr: nil,
		},
		{
			name: "test starts with word",
			fn: func() (bool, error) {
				rootNode := ds.TrieNode{
					Children: map[rune]*ds.TrieNode{},
				}
				trie := ds.Trie{Root: &rootNode}
				trie.Insert("test")
				found := trie.StartsWith("te")
				return found, nil
			},
			expectedRes: true,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		res, err := test.fn()
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedRes, res)
	}
}
