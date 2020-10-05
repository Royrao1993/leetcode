package design

type Trie struct {
	Next  [26]*Trie
	IsEnd bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		if node.Next[v-'a'] == nil {
			node.Next[v-'a'] = NewTrie()
		}
		node = node.Next[v-'a']
	}
	node.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word {
		if node.Next[v-'a'] == nil {
			return false
		}
		node = node.Next[v-'a']
	}
	return node.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, v := range prefix {
		if node.Next[v-'a'] == nil {
			return false
		}
		node = node.Next[v-'a']
	}
	return true
}
