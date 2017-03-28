package tree

// Root the root of a tree or sub-tree
type Root interface {
	Value() interface{}
	SetValue(interface{})

	Childrens() []Child
	Add(...Child)
	Remove(...Child)
	Has(...Child)
	Each(func(children Child) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Root) bool
}

// Child describes the children of a root.
// Child is also the root of its own sub-tree.
type Child interface {
	Root
	Parent() Root
}

// Interface describes a Tree
type Interface interface {
	Root() Root
	WalkD(func(node Root) bool)
	WalkB(func(node Root) bool)
	Has(...Root)

	Values() []interface{}
	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool
	CopyTree() Interface
}
