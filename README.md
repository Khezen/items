[![Build Status](http://img.shields.io/travis/Khezen/array.svg?style=flat-square)](https://travis-ci.org/Khezen/items) [![codecov](https://img.shields.io/codecov/c/github/Khezen/array/master.svg?style=flat-square)](https://codecov.io/gh/Khezen/items)
[![Go Report Card](https://goreportcard.com/badge/github.com/khezen/array?style=flat-square)](https://goreportcard.com/report/github.com/khezen/items)

# Base collection [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/items/collection)

`
import "github.com/khezen/items/collection"
`

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

# Array [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/items/array)

`
import "github.com/khezen/items/array"
`

Abstraction layer over slices exposing utility functions and thread safe implementation of dynamic array.

```golang
type Interface interface {
	collection.Interface
	Get(i int) (interface{}, error)
	Insert(i int, item ...interface{}) error
	RemoveAt(i int) (interface{}, error)
	ReplaceAt(i int, substitute interface{}) (interface{}, error)
	IndexOf(interface{}) (int, error)
	Swap(i, j int) error

	Copy() Interface
}
```

```golang
package example

import "github.com/khezen/items/array"

arr := array.New(0, 2, -4, 10)
threadsafeArr := array.NewTS(0, 2, -4, 10)
```


# Set [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/items/set)

`
import "github.com/khezen/items/set"
`

Both threadsafe and non-threadsafe implementations of a generic
set data structure.

```Golang
type Interface interface {
	Add(items ...interface{})
	Remove(items ...interface{})
	Pop() interface{}
	Has(items ...interface{}) bool
	Each(func(interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(s Interface) bool
	IsSubset(s Interface) bool
	IsSuperset(s Interface) bool
	Merge(s Interface)
	Separate(s Interface)
	Retain(s Interface)

	String() string
	Slice() []interface{}
	Copy() Interface
}
```

```golang
package example

import "github.com/khezen/items/set"

set := set.New(0, 2, -4, 10)
threadsafeSet := set.NewTS(0, 2, -4, 10)
```
