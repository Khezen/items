[![Build Status](http://img.shields.io/travis/Khezen/array.svg?style=flat-square)](https://travis-ci.org/Khezen/items) [![codecov](https://img.shields.io/codecov/c/github/Khezen/array/master.svg?style=flat-square)](https://codecov.io/gh/Khezen/items)
[![Go Report Card](https://goreportcard.com/badge/github.com/khezen/array?style=flat-square)](https://goreportcard.com/report/github.com/khezen/items)

# Collections for Go

## Array [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/items/array)

`
import "github.com/khezen/items/array"
`

Dynamic array data structure for Go.

Abstraction layer over slices exposing utility functions and thread safe implementation of dynamic array.

```golang
type Interface interface {
  Get(i int) (interface{}, error)
  Add(items ...interface{})
  Insert(i int, item ...interface{}) error
  Remove(items ...interface{})
  RemoveAt(i int) (interface{}, error)
  Replace(toBeReplaced, substitute interface{})
  ReplaceAt(i int, substitute interface{}) (interface{}, error)
  IndexOf(interface{}) (int, error)
  Swap(i, j int) error
  Has(items ...interface{}) bool

  Each(func(interface{}) bool)
  Len() int
  Clear()
  IsEmpty() bool

  Merge(a Interface)
  Separate(a Interface)
  Retain(a Interface)

  String() string
  Slice() []interface{}
  Copy() Interface
}
```

```golang
func Union(arrays ...Interface) Interface
```
```golang
func Difference(arrays ...Interface) Interface
```
```golang
func Intersection(arrays ...Interface) Interface
```
```golang
func Exclusion(arrays ...Interface) Interface
```

#### Example

```golang
package example

import "github.com/khezen/items/array"

arr := array.New(0, 2, -4, 10)
threadsafeArr := array.NewTS(0, 2, -4, 10)
```


## Set [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/items/set)

`
import "github.com/khezen/items/set"
`


set is a basic and simple, hash-based, **set** data structure implementation
in Go (Golang).

set provides both threadsafe and non-threadsafe implementations of a generic
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
func Union(set1, set2 Interface, sets ...Interface) Interface
```
```golang
func Difference(set1, set2 Interface, sets ...Interface) Interface
```
```golang
func Intersection(set1, set2 Interface, sets ...Interface) Interface
```
```golang
func Exclusion(set1 Interface, set2 Interface, sets ...Interface) Interface
```


#### Example
```golang
package example

import "github.com/khezen/items/set"

set := set.New(0, 2, -4, 10)
threadsafeSet := set.NewTS(0, 2, -4, 10)
```
