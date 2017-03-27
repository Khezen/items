package oset

import (
	"github.com/khezen/struct/array"
	"github.com/khezen/struct/set"
)

type oset struct {
	a array.Interface
	s set.Interface
}

func New(items ...interface{}) Interface {
	s := &oset{
		array.New(),
		set.New(items...),
	}
	s.s.Each(func(item interface{}) bool {
		s.a.Add(item)
		return true
	})
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

func (s *oset) Remove(items ...interface{}) {
	s.a.Remove(items...)
	s.s.Remove(items...)
}

func (s *oset) Replace(item, substitute interface{}) {
	s.a.Replace(item, substitute)
	s.s.Replace(item, substitute)
}

func (s *oset) Has(items ...interface{}) bool {
	return s.s.Has(items...)
}

func (s *oset) Each(f func(item interface{}) bool) {
	s.a.Each(f)
}
