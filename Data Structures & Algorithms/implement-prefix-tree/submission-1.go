type PrefixTreeNode struct {
	children [26]*PrefixTreeNode
	isEnd bool
}

func NewPrefixTreeNode() *PrefixTreeNode {
	return &PrefixTreeNode{}
}
type PrefixTree struct {
	root *PrefixTreeNode

}

func Constructor() PrefixTree {
	return PrefixTree{
		root: NewPrefixTreeNode(),
	}
    
}

func (this *PrefixTree) Insert(word string) {
	node := this.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = NewPrefixTreeNode()
		}
		node = node.children[index]
	}
	node.isEnd = true

}

func (this *PrefixTree) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd

}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true

}
