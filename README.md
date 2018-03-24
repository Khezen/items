[![Build Status](http://img.shields.io/travis/khezen/struct/master.svg?style=flat-square)](https://travis-ci.org/khezen/struct) [![codecov](https://img.shields.io/codecov/c/github/khezen/struct/master.svg?style=flat-square)](https://codecov.io/gh/khezen/struct)

# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/collection) *collection*

`
import "github.com/khezen/struct/collection"
`

Exposes base collection interface and mixing operations(union, intersection, etc...)

```golang
type Interface interface {
	Add(...interface{})
	Remove(...interface{})
	Replace(item, substitute interface{})
	Has(...interface{}) bool
	Each(func(item interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool

	Merge(Interface)
	Separate(Interface)
	Retain(Interface)

	String() string
	Slice() []interface{}
	CopyCollection() Interface
}
```

```golang
func Union(collections ...Interface) Interface
```
```golang
func Difference(collections ...Interface) Interface
```
```golang
func Intersection(collections ...Interface) Interface
```
```golang
func Exclusion(collections ...Interface) Interface
```

# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/array) *array*

`
import "github.com/khezen/struct/array"
`

Abstraction layer over slices exposing utility functions and synchronized implementation of dynamic array.



# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/set) *set*

`
import "github.com/khezen/struct/set"
`

Both synchronized and non-synchronized implementations of a generic set data structure.


# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/oset) *ordered set*

`
import "github.com/khezen/struct/oset"
`

Both synchronized and non-synchronized implementations of a generic ordered set data structure.


# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/hashmap) *hashmap*

`
import "github.com/khezen/struct/hashmap"
`

Both synchronized and non-synchronized implementations of a generic
hashmap data structure.
