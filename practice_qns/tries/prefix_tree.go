package tries

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: map[rune]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		if _, found := curr.children[c]; !found {
			trie := &Trie{
				children: map[rune]*Trie{},
			}
			curr.children[c] = trie
		}

		curr = curr.children[c]
	}

	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		if _, found := curr.children[c]; !found {
			return false
		}
		curr = curr.children[c]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		if _, found := curr.children[c]; !found {
			return false
		}
		curr = curr.children[c]
	}
	return true
}
