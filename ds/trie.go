package ds

type TrieNode struct {
	Children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(word string) {
	curr := t.Root
	for _, ch := range word {
		if _, found := curr.Children[ch]; !found {
			newNode := &TrieNode{Children: map[rune]*TrieNode{}}
			curr.Children[ch] = newNode
		}
		curr = curr.Children[ch]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.Root
	for _, ch := range word {
		if _, found := curr.Children[ch]; !found {
			return false
		}
		curr = curr.Children[ch]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(word string) bool {
	curr := t.Root
	for _, ch := range word {
		if _, found := curr.Children[ch]; !found {
			return false
		}
		curr = curr.Children[ch]
	}
	return true
}
