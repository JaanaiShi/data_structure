package radix1


type Node struct {
	Val string
	Children map[string]*Node
	IsEnd bool
}

type Tree struct {
	Node *Node
}

const (
	RootVal = "#"
)



func NewRadixTree() *Node {
	return &Node{
		Val: RootVal,
		Children: make(map[string]*Node),
		IsEnd: true,
	}
}


func (t *Node) Insert(str string) {
	
}
