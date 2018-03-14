package oset

import (
	"github.com/khezen/struct/array"
	"github.com/khezen/struct/collection"
	"github.com/khezen/struct/set"
)

type oset struct {
	a array.Interface
	s set.Interface
}

// New creates a new ordered set
func New(items ...interface{}) Interface {
	s := &oset{
		array.New(),
		set.New(),
	}
	for _, item := range items {
		if !s.Has(item) {
			s.s.Add(item)
			s.a.Add(item)
		}
	}
	return s
}

func (s *oset) Get(i int) (interface{}, error) {
	return s.a.Get(i)
}

func (s *oset) Add(items ...interface{}) {
	for _, item := range items {
		if !s.Has(item) {
			s.a.Add(item)
			s.s.Add(item)
		}
	}
}

func (s *oset) Insert(i int, items ...interface{}) error {
	toInsert := make([]interface{}, 0, len(items))
	for _, item := range items {
		if !s.Has(item) {
			toInsert = append(toInsert, item)
		}
	}
	err := s.a.Insert(i, toInsert...)
	if err != nil {
		return err
	}
	s.s.Add(toInsert...)
	return nil
}

func (s *oset) Remove(items ...interface{}) {
	s.a.Remove(items...)
	s.s.Remove(items...)
}

func (s *oset) RemoveAt(i int) (interface{}, error) {
	item, err := s.a.RemoveAt(i)
	if err != nil {
		return nil, err
	}
	s.s.Remove(item)
	return item, nil
}

func (s *oset) Replace(item, substitute interface{}) {
	s.a.Replace(item, substitute)
	s.s.Replace(item, substitute)
}

func (s *oset) ReplaceAt(i int, substitute interface{}) (interface{}, error) {
	item, err := s.a.ReplaceAt(i, substitute)
	if err != nil {
		return nil, err
	}
	s.s.Replace(item, substitute)
	return item, nil
}

func (s *oset) IndexOf(item interface{}) (int, error) {
	return s.a.IndexOf(item)
}

func (s *oset) Swap(i, j int) {
	s.a.Swap(i, j)
}

func (s *oset) Has(items ...interface{}) bool {
	return s.s.Has(items...)
}

func (s *oset) Each(f func(item interface{}) bool) {
	s.a.Each(f)
}

func (s *oset) Len() int {
	return s.a.Len()
}

func (s *oset) Clear() {
	s.a.Clear()
	s.s.Clear()
}

func (s *oset) IsEmpty() bool {
	return s.a.IsEmpty()
}

func (s *oset) IsEqual(t collection.Interface) bool {
	if conv, ok := t.(*osetSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*osetSortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	return s.a.IsEqual(t)
}

func (s *oset) IsSubset(t Interface) bool {
	if conv, ok := t.(*osetSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*osetSortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	return s.s.IsSubset(t.Set())
}

func (s *oset) IsSuperset(t Interface) bool {
	if conv, ok := t.(*osetSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*osetSortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	return s.s.IsSuperset(t.Set())
}

func (s *oset) Merge(t collection.Interface) {
	if conv, ok := t.(*osetSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*osetSortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	s.s.Merge(t)
	s.a.Merge(t)
}

func (s *oset) Separate(t collection.Interface) {
	if conv, ok := t.(*osetSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*osetSortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	s.s.Separate(t)
	s.a.Separate(t)
}

func (s *oset) Retain(t collection.Interface) {
	if conv, ok := t.(*osetSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	if conv, ok := t.(*osetSortSync); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	s.s.Retain(t)
	s.a.Retain(t)
}

func (s *oset) SubArray(i, j int) (array.Interface, error) {
	return s.a.SubArray(i, j)
}

func (s *oset) Subset(i, j int) (Interface, error) {
	arr, err := s.SubArray(i, j)
	if err != nil {
		return nil, err
	}
	sub := New(arr.Slice()...)
	return sub, nil
}

func (s *oset) String() string {
	return s.a.String()
}

func (s *oset) Slice() []interface{} {
	return s.a.Slice()
}

func (s *oset) CopyOset() Interface {
	return New(s.Slice()...)
}

func (s *oset) CopyArr() array.Interface {
	return s.a.CopyArr()
}

func (s *oset) CopySet() set.Interface {
	return s.s.CopySet()
}

func (s *oset) CopyCollection() collection.Interface {
	return s.CopyOset()
}

func (s *oset) Arr() array.Interface {
	return s.a
}

func (s *oset) Set() set.Interface {
	return s.s
}
