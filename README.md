[![Build Status](http://img.shields.io/travis/Khezen/struct.svg?style=flat-square)](https://travis-ci.org/Khezen/struct) [![codecov](https://img.shields.io/codecov/c/github/Khezen/struct/master.svg?style=flat-square)](https://codecov.io/gh/Khezen/struct)

# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/collection) Collection

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

# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/array) Array

`
import "github.com/khezen/struct/array"
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

	SubArray(i, j int) (Interface, error)
	Copy() Interface
}

```

```golang
package example

import "github.com/khezen/struct/array"

arr := array.New(0, 2, -4, 10)
threadsafeArr := array.NewTS(0, 2, -4, 10)
```


# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/set) Set

`
import "github.com/khezen/struct/set"
`

Both threadsafe and non-threadsafe implementations of a generic
set data structure.

```Golang
type Interface interface {
	collection.Interface
	Pop() interface{}
	IsSubset(s Interface) bool
	IsSuperset(s Interface) bool
	Copy() Interface
}
```

```golang
package example

import "github.com/khezen/struct/set"

s := set.New(0, 2, -4, 10)
threadsafeS := set.NewTS(0, 2, -4, 10)
```

# [![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/khezen/struct/hashmap) Hashmap

`
import "github.com/khezen/struct/hashmap"
`

Both threadsafe and non-threadsafe implementations of a generic
hashmap data structure.

```Golang
type Interface interface {
	Get(k interface{}) (interface{}, error)
	Put(k, v interface{})
	Remove(keys ...interface{})
	Has(keys ...interface{}) bool
	HasValue(values ...interface{}) bool
	KeyOf(value interface{}) (interface{}, error)
	Each(func(k, v interface{}) bool)

	Len() int
	Clear()
	IsEmpty() bool
	IsEqual(Interface) bool

	String() string
	Keys() []interface{}
	Values() []interface{}
	Copy() Interface
}
```

```golang
package example

import "github.com/khezen/struct/hashmap"

hm := hashmap.New()
threadsafeHm := hashmap.NewTS()
```
