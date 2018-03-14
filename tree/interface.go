package tree

// Node has children
type Node interface {
	Children() []Node
	Add(...Node)
	Remove(...Node)
	Has(...Node)
	Each(func(child Node) bool)
	Len() int
	Clear()
	IsLeaf() bool
	IsEqual(Node) bool
}

// ValuedNode the root of a tree or sub-tree
type ValuedNode interface {
	Node
	Value() interface{}
	SetValue(interface{})
}

// Interface describes a Tree
type Interface interface {
	Root() Node
	Walk(func(node Node) bool)
	Has(...Node)
	Nodes() []Node
	Values() []interface{}
	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool
	CopyTree() Interface
}
