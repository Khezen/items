package tree

// Node the root of a tree or sub-tree
type Node interface {
	Value() interface{}
	SetValue(interface{})
	Children() []Node
	Add(...Node)
	Remove(...Node)
	Has(...Node)
	Each(func(child Node) bool)
	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Node) bool
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
