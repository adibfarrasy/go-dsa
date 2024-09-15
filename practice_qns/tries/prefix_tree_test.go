package tries_test

import (
	"go-data-structure/practice_qns/tries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixTree(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			trie := tries.Constructor()
			trie.Insert("apple")

			got := trie.Search("apple")
			assert.Equal(t, true, got)
			got = trie.Search("app")
			assert.Equal(t, false, got)
			got = trie.StartsWith("app")
			assert.Equal(t, true, got)

			trie.Insert("app")
			got = trie.Search("app")
			assert.Equal(t, true, got)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
