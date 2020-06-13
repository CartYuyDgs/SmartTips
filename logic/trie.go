package logic

type Node struct {
	char   rune
	childs map[rune]*Node
	Date   interface{}
	deep   int
	isTerm bool
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
	var parent *Node = t.root
	allChars := []rune(key)
	for _, c := range allChars {
		node, ok := parent.childs[c]
		if !ok {
			node = NewNode(c, parent.deep+1)
			parent.childs[c] = node
		}

		parent = node
	}

	parent.Date = data
	parent.isTerm = true
}

func (t *Trie) PrefixSearch(key string) (nodes []*Node) {
	var parent = t.root
	allChars := []rune(key)
	for _, c := range allChars {
		node, ok := parent.childs[c]
		if !ok {
			return
		}

		parent = node
	}
	var queue []*Node
	queue = append(queue, parent)
	for len(queue) > 0 {
		var q2 []*Node
		for _, n := range queue {
			if n.isTerm == true {
				nodes = append(nodes, n)
				continue
			}

			for _, v := range n.childs {
				q2 = append(q2, v)
			}
		}
		queue = q2
	}

	return
}
