package collection

// Interface describes method exposed by a collection
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

// Union is the merger of multiple collections. It returns a new collection with all the
// elements present in all the collections that are passed.
//
// The dynamic type of the returned collection is determined by the first passed collection's
// implementation of the New() method.
func Union(collections ...Interface) Interface {
	if len(collections) == 0 {
		return nil
	}

	u := collections[0].CopyCollection()
	collections = collections[1:]
	for _, collection := range collections {
		collection.Each(func(item interface{}) bool {
			if !u.Has(item) {
				u.Add(item)
			}
			return true
		})
	}
	return u
}

// Difference returns a new collection which contains items which are in in the first
// collection but not in the others. Unlike the Difference() method you can use this
// function separately with multiple collections.
func Difference(collections ...Interface) Interface {
	if len(collections) == 0 {
		return nil
	}

	s := collections[0].CopyCollection()
	collections = collections[1:]
	for _, collection := range collections {
		s.Separate(collection) // separate is thread safe
	}
	return s
}

// Intersection returns a new collection which contains items that only exist in all given collections.
func Intersection(collections ...Interface) Interface {
	if len(collections) == 0 {
		return nil
	}
	result := Union(collections...)
	for _, collection := range collections {
		result.Retain(collection)
	}
	return result
}

// Exclusion returns a new collection which s is the difference of items which are in
// one of either, but not in both.
func Exclusion(collections ...Interface) Interface {
	length := len(collections)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return collections[0]
	}
	intersections := make([]Interface, 0, length)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if j != i {
				intersections = append(intersections, Intersection(collections[i], collections[j]))
			}
		}
	}
	intersection := Union(intersections...)
	exclusion := Union(collections...)
	exclusion.Separate(intersection)
	return exclusion
}
