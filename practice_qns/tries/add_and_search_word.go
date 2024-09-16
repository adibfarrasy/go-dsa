package tries

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func NewWD() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: map[byte]*TrieNode{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i := range word {
		if _, found := curr.children[word[i]]; !found {
			curr.children[word[i]] = &TrieNode{
				children: map[byte]*TrieNode{},
			}
		}
		curr = curr.children[word[i]]
	}

	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(j int, root *TrieNode) bool {
		cur := root

		for i := j; i < len(word); i++ {
			c := word[i]
			if c == '.' {
				for _, child := range cur.children {
					if dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := cur.children[c]; !ok {
					return false
				}

				cur = cur.children[c]
			}
		}

		return cur.isEnd
	}

	return dfs(0, this.root)
}
