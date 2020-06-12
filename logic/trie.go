package logic

type Node struct {
	char   rune
	childs map[rune]*Node
	Date   interface{}
	deep   int
}

type Trie struct {
	root *Node
	size int
}

func NewNode(char rune, deep int) *Node {
	return &Node{
		char:   char,
		childs: make(map[rune]*Node, 16),
		Date:   nil,
		deep:   deep,
	}
}

func NewTrie() *Trie {
	trie := &Trie{
		root: NewNode(' ', 1),
		size: 0,
	}
	return trie
}

func (t *Trie) Add(key string, data interface{}) {

}
