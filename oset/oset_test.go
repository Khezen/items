package oset

import (
	"testing"

	"github.com/khezen/struct/array"
	"github.com/khezen/struct/collection"
	"github.com/khezen/struct/set"
)

func testErr(err error, expectErr bool, t *testing.T) {
	if (expectErr && err == nil) || (!expectErr && err != nil) {
		t.Errorf(" Error expected? %v. Got: %v.", expectErr, err)
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		oset      Interface
		i         int
		expected  interface{}
		expectErr bool
	}{
		{New(1, 7, -5), 2, -5, false},
		{New(1, 7, -5), -1, nil, true},
		{New(1, 7, -5), 1000, nil, true},
		{NewSync(1, 7, -5), 2, -5, false},
		{NewSync(1, 7, -5), -1, nil, true},
		{NewSync(1, 7, -5), 1000, nil, true},
	}
	for _, c := range cases {
		item, err := c.oset.Get(c.i)
		testErr(err, c.expectErr, t)
		if item != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, item)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		oset, toBeAdded, expected Interface
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, -8, 42, -1)},
		{New(), New(42, -1), New(42, -1)},
		{New(), New(), New()},
		{New(), New(nil), New(nil)},
		{New(), New(42, -1), New(42, -1)},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8, 42, -1)},
		{NewSync(), New(42, -1), NewSync(42, -1)},
		{NewSync(), New(), NewSync()},
		{NewSync(), New(nil), NewSync(nil)},
		{NewSync(), New(42, -1), NewSync(42, -1)},
	}
	for _, c := range cases {
		c.oset.Add(c.toBeAdded.Slice()...)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		oset, toBeInserted, expected Interface
		i                            int
		expectErr                    bool
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, 42, -1, -8), 2, false},
		{New(1, 4, -8), New(42, -1), New(1, 4, -8), -1, true},
		{New(1, 4, -8), New(42, -1), New(1, 4, -8), 42, true},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, 42, -1, -8), 2, false},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8), -1, true},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8), 42, true},
	}
	for _, c := range cases {
		err := c.oset.Insert(c.i, c.toBeInserted.Slice()...)
		testErr(err, c.expectErr, t)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		oset, toBeRemoved, expected Interface
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, -8)},
		{New(1, 4, -8), New(1, -8), New(4)},
		{New(), New(42, -1), New()},
		{New(), New(), New()},
		{New(), New(nil), New()},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8)},
		{NewSync(1, 4, -8), New(1, -8), NewSync(4)},
		{NewSync(), New(42, -1), NewSync()},
		{NewSync(), New(), NewSync()},
		{NewSync(), New(nil), NewSync()},
	}
	for _, c := range cases {
		c.oset.Remove(c.toBeRemoved.Slice()...)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestRemoveAt(t *testing.T) {
	cases := []struct {
		oset, expected Interface
		i              int
		removed        interface{}
		expectErr      bool
	}{
		{New(1, 4, -8), New(4, -8), 0, 1, false},
		{New(1, 4, -8), New(1, 4, -8), -1, nil, true},
		{New(1, 4, -8), New(1, 4, -8), 42, nil, true},
		{NewSync(1, 4, -8), New(4, -8), 0, 1, false},
		{NewSync(1, 4, -8), New(1, 4, -8), -1, nil, true},
		{NewSync(1, 4, -8), New(1, 4, -8), 42, nil, true},
	}
	for _, c := range cases {
		removed, err := c.oset.RemoveAt(c.i)
		testErr(err, c.expectErr, t)
		if removed != c.removed {
			t.Errorf("Expected %v. Got %v.", c.removed, removed)
		}
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestReplace(t *testing.T) {
	cases := []struct {
		oset, expected   Interface
		item, substitute interface{}
	}{
		{New(1, 4, -8), New(42, 4, -8), 1, 42},
		{New(1, 4, -8), New(1, 4, -8), 1000, 42},
		{NewSync(1, 4, -8), NewSync(42, 4, -8), 1, 42},
		{NewSync(1, 4, -8), NewSync(1, 4, -8), 1000, 42},
	}
	for _, c := range cases {
		c.oset.Replace(c.item, c.substitute)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestReplaceAt(t *testing.T) {
	cases := []struct {
		oset, expected Interface
		i              int
		substitute     interface{}
	}{
		{New(1, 4, -8), New(1, 42, -8), 1, 42},
		{New(1, 4, -8), New(1, 4, -8), 1000, 42},
		{NewSync(1, 4, -8), NewSync(1, 42, -8), 1, 42},
		{NewSync(1, 4, -8), NewSync(1, 4, -8), 1000, 42},
	}
	for _, c := range cases {
		c.oset.ReplaceAt(c.i, c.substitute)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.oset.Slice())
		}
	}

}

func TestIndexOf(t *testing.T) {
	cases := []struct {
		oset      Interface
		item      interface{}
		i         int
		expectErr bool
	}{
		{New(1, 42, -8), 42, 1, false},
		{New(1, 42, -8), 1000, -1, true},
		{NewSync(1, 42, -8), 42, 1, false},
		{NewSync(1, 42, -8), 1000, -1, true},
	}
	for _, c := range cases {
		i, err := c.oset.IndexOf(c.item)
		testErr(err, c.expectErr, t)
		if i != c.i {
			t.Errorf("Expected %v. Got %v.", c.i, i)
		}
	}
}

func TestSubArray(t *testing.T) {
	cases := []struct {
		oset, expected Interface
		i, j           int
		expectErr      bool
	}{
		{New(1, 42, -8, 12), New(42, -8), 1, 2, false},
		{New(1, 42, -8, 12), nil, -1, 2, true},
		{New(1, 42, -8, 12), nil, 1000, 2, true},
		{New(1, 42, -8, 12), nil, 1, -2, true},
		{New(1, 42, -8, 12), nil, 1, 1000, true},
		{New(1, 42, -8, 12), nil, 2, 1, true},
		{NewSync(1, 42, -8, 12), NewSync(42, -8), 1, 2, false},
		{NewSync(1, 42, -8, 12), nil, -1, 2, true},
		{NewSync(1, 42, -8, 12), nil, 1000, 2, true},
		{NewSync(1, 42, -8, 12), nil, 1, -2, true},
		{NewSync(1, 42, -8, 12), nil, 1, 1000, true},
		{NewSync(1, 42, -8, 12), nil, 2, 1, true},
	}
	for _, c := range cases {
		arr, err := c.oset.SubArray(c.i, c.j)
		testErr(err, c.expectErr, t)
		if !c.expectErr {
			if !arr.IsEqual(c.expected) {
				t.Errorf("Expected %v. Got %v.", c.expected.Slice(), arr.Slice())
			}
			if arr.IsEqual(c.oset) {
				t.Errorf("c.oset should not be modified")
			}
		}
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		oset, expected Interface
		i, j           int
		expectErr      bool
	}{
		{New(1, 42, -8), New(42, 1, -8), 0, 1, false},
		{New(1, 42, -8), New(42, 1, -8), 1, 0, false},
		{New(1, 42, -8), New(1, 42, -8), -1, 0, true},
		{New(1, 42, -8), New(1, 42, -8), 1000, 0, true},
		{New(1, 42, -8), New(1, 42, -8), 1, 1000, true},
		{New(1, 42, -8), New(1, 42, -8), 1, -1, true},
		{NewSync(1, 42, -8), NewSync(42, 1, -8), 0, 1, false},
		{NewSync(1, 42, -8), NewSync(42, 1, -8), 1, 0, false},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), -1, 0, true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), 1000, 0, true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), 1, 1000, true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), 1, -1, true},
	}
	for _, c := range cases {
		c.oset.Swap(c.i, c.j)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestHas(t *testing.T) {
	cases := []struct {
		oset, items Interface
		expected    bool
	}{
		{New(1, 42, -8), New(1, 42, -8), true},
		{New(1, 42, -8), New(-8), true},
		{New(1, 42, -8), New(34), false},
		{New(1, 42, -8), New(nil), false},
		{New(1, 42, -8), New(), true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), true},
		{NewSync(1, 42, -8), NewSync(-8), true},
		{NewSync(1, 42, -8), NewSync(34), false},
		{NewSync(1, 42, -8), NewSync(nil), false},
		{NewSync(1, 42, -8), NewSync(), true},
	}
	for _, c := range cases {
		has := c.oset.Has(c.items.Slice()...)
		if has != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, has)
		}
	}
}

func TestEach(t *testing.T) {
	cases := []struct {
		oset              Interface
		counter, expected int
	}{
		{New(1, -8, 42), 0, 2},
		{NewSync(1, -8, 42), 0, 2},
	}
	for _, c := range cases {
		callback := func(item interface{}) bool {
			value := item.(int)
			c.counter++
			return value > 0
		}
		c.oset.Each(callback)
		if c.counter != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, c.counter)
		}
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		oset Interface
		len  int
	}{
		{New(), 0},
		{New(1), 1},
		{New(1, 42, -8), 3},
		{NewSync(), 0},
		{NewSync(1), 1},
		{NewSync(1, 42, -8), 3},
	}
	for _, c := range cases {
		if c.oset.Len() != c.len {
			t.Errorf("Expected %v. Got %v.", c.len, c.oset.Len())
		}
	}
}

func TestClear(t *testing.T) {
	cases := []struct {
		oset Interface
	}{
		{New(1, 42, -8)},
		{New()},
		{NewSync(1, 42, -8)},
		{NewSync()},
	}
	for _, c := range cases {
		c.oset.Clear()
		if !c.oset.IsEmpty() {
			t.Error("Array should be empty")
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		oset    Interface
		isEmpty bool
	}{
		{New(), true},
		{New(1, 42, -8), false},
		{NewSync(), true},
		{NewSync(1, 42, -8), false},
	}
	for _, c := range cases {
		if c.oset.IsEmpty() != c.isEmpty {
			t.Errorf("Expected %v. Got %v.", c.isEmpty, c.oset.IsEmpty())
		}
	}
}

func TestIsEqual(t *testing.T) {
	cases := []struct {
		oset, toBeCompared Interface
		isEqual            bool
	}{
		{New(), New(), true},
		{New(1, 42, -8), New(1, 42, -8), true},
		{New(1, 42, -8), New(1, "42", -8), false},
		{New(1, 42, -8), New(), false},
		{New(1, 42, -8), New(42, 1, -8), false},
		{New(66, -1000), New(42, 1, 8), false},
		{NewSync(), NewSync(), true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), true},
		{NewSync(1, 42, -8), NewSync(1, "42", -8), false},
		{NewSync(1, 42, -8), NewSync(), false},
		{NewSync(1, 42, -8), NewSync(42, 1, -8), false},
		{NewSync(66, -1000), NewSync(42, 1, 8), false},
	}
	for _, c := range cases {
		isEqual := c.oset.IsEqual(c.toBeCompared)
		if isEqual != c.isEqual {
			t.Errorf("Expected %v to be equal to %v? %v. Got: %v", c.oset.Slice(), c.toBeCompared.Slice(), c.isEqual, isEqual)
		}
	}
}

func TestIsSubset(t *testing.T) {
	cases := []struct {
		s, sub Interface
		isSub  bool
	}{
		{New("1", "2", "3", "4"), New("1", "2", "3"), true},
		{New("1", "2", "3"), New("1", "2", "3", "4"), false},
		{NewSync("1", "2", "3", "4"), NewSync("1", "2", "3"), true},
		{NewSync("1", "2", "3"), NewSync("1", "2", "3", "4"), false},
		{NewSortedSync(nil, "1", "2", "3", "4"), NewSortedSync(nil, "1", "2", "3"), true},
		{NewSortedSync(nil, "1", "2", "3"), NewSortedSync(nil, "1", "2", "3", "4"), false},
	}
	for _, c := range cases {
		ok := c.s.IsSubset(c.sub)
		if ok != c.isSub {
			t.Errorf("Expected %v. Got %v", c.isSub, ok)
		}
	}
}

func TestIsSuperset(t *testing.T) {
	cases := []struct {
		s, sub Interface
		isSub  bool
	}{
		{New("1", "2", "3", "4"), New("1", "2", "3"), false},
		{New("1", "2", "3"), New("1", "2", "3", "4"), true},
		{NewSync("1", "2", "3", "4"), NewSync("1", "2", "3"), false},
		{NewSync("1", "2", "3"), NewSync("1", "2", "3", "4"), true},
		{NewSortedSync(nil, "1", "2", "3", "4"), NewSortedSync(nil, "1", "2", "3"), false},
		{NewSortedSync(nil, "1", "2", "3"), NewSortedSync(nil, "1", "2", "3", "4"), true},
	}
	for _, c := range cases {
		ok := c.s.IsSuperset(c.sub)
		if ok != c.isSub {
			t.Errorf("Expected %v. Got %v", c.isSub, ok)
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		oset, toBeMerged, expected Interface
	}{
		{New(1, 42), New(-8), New(1, 42, -8)},
		{New(1, 42), New(-8, nil), New(1, 42, -8, nil)},
		{New(1, 42), New(), New(1, 42)},
		{New(), New(), New()},
		{NewSync(1, 42), NewSync(-8), NewSync(1, 42, -8)},
		{NewSync(1, 42), NewSync(-8, nil), NewSync(1, 42, -8, nil)},
		{NewSync(1, 42), NewSync(), NewSync(1, 42)},
		{NewSync(), NewSync(), NewSync()},
		{NewSortedSync(nil, 1, 42), NewSortedSync(nil, -8), NewSortedSync(nil, 1, 42, -8)},
		{NewSortedSync(nil, 1, 42), NewSortedSync(nil, -8, nil), NewSortedSync(nil, 1, 42, -8, nil)},
		{NewSortedSync(nil, 1, 42), NewSortedSync(nil), NewSortedSync(nil, 1, 42)},
		{NewSortedSync(nil), NewSortedSync(nil), NewSortedSync(nil)},
	}
	for _, c := range cases {
		c.oset.Merge(c.toBeMerged)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestSeparate(t *testing.T) {
	cases := []struct {
		oset, toBeMerged, expected Interface
	}{
		{New(1, 42, -8), New(1, 42), New(-8)},
		{New(1, 42, -8), New(1, 42, nil), New(-8)},
		{New(1, 42, -8), New(), New(1, 42, -8)},
		{New(), New(), New()},
		{NewSync(1, 42, -8), NewSync(1, 42), NewSync(-8)},
		{NewSync(1, 42, -8), NewSync(1, 42, nil), NewSync(-8)},
		{NewSync(1, 42, -8), NewSync(), NewSync(1, 42, -8)},
		{NewSync(), NewSync(), NewSync()},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 1, 42), NewSortedSync(nil, -8)},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 1, 42, nil), NewSortedSync(nil, -8)},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil), NewSortedSync(nil, 1, 42, -8)},
		{NewSortedSync(nil), NewSortedSync(nil), NewSortedSync(nil)},
	}
	for _, c := range cases {
		c.oset.Separate(c.toBeMerged)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestRetain(t *testing.T) {
	cases := []struct {
		oset, toBeMerged, expected Interface
	}{
		{New(1, 42, -8), New(1, -8, 100), New(1, -8)},
		{New(1, 42, -8), New(1, -8, 100, nil), New(1, -8)},
		{New(1, 42, -8), New(), New()},
		{New(), New(), New()},
		{NewSync(1, 42, -8), NewSync(1, -8, 100), NewSync(1, -8)},
		{NewSync(1, 42, -8), NewSync(1, -8, 100, nil), NewSync(1, -8)},
		{NewSync(1, 42, -8), NewSync(), NewSync()},
		{NewSync(), NewSync(), NewSync()},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 1, -8, 100), NewSortedSync(nil, 1, -8)},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 1, -8, 100, nil), NewSortedSync(nil, 1, -8)},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil), NewSortedSync(nil)},
		{NewSortedSync(nil), NewSortedSync(nil), NewSortedSync(nil)},
	}
	for _, c := range cases {
		c.oset.Retain(c.toBeMerged)
		if !c.oset.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.oset.Slice())
		}
	}
}

func TestSubset(t *testing.T) {
	cases := []struct {
		s, expected Interface
		i, j        int
		expectErr   bool
	}{
		{New(1, 42, -8, 12), New(42, -8), 1, 2, false},
		{New(1, 42, -8, 12), nil, -1, 2, true},
		{New(1, 42, -8, 12), nil, 1000, 2, true},
		{New(1, 42, -8, 12), nil, 1, -2, true},
		{New(1, 42, -8, 12), nil, 1, 1000, true},
		{New(1, 42, -8, 12), nil, 2, 1, true},
		{NewSync(1, 42, -8, 12), NewSync(42, -8), 1, 2, false},
		{NewSync(1, 42, -8, 12), nil, -1, 2, true},
		{NewSync(1, 42, -8, 12), nil, 1000, 2, true},
		{NewSync(1, 42, -8, 12), nil, 1, -2, true},
		{NewSync(1, 42, -8, 12), nil, 1, 1000, true},
		{NewSync(1, 42, -8, 12), nil, 2, 1, true},
	}
	for _, c := range cases {
		s, err := c.s.Subset(c.i, c.j)
		testErr(err, c.expectErr, t)
		if !c.expectErr {
			if !s.IsEqual(c.expected) {
				t.Errorf("Expected %v. Got %v.", c.expected.Slice(), s.Slice())
			}
			if s.IsEqual(c.s) {
				t.Errorf("c.array should not be modified")
			}
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		oset     Interface
		expected string
	}{
		{New(1, 2, 3), "[1 2 3]"},
		{New(-12, 6, 111), "[-12 6 111]"},
		{New(), "[]"},
		{New(nil), "[<nil>]"},
		{NewSync(1, 2, 3), "[1 2 3]"},
		{NewSync(-12, 6, 111), "[-12 6 111]"},
		{NewSync(), "[]"},
		{NewSync(nil), "[<nil>]"},
	}

	for _, c := range cases {
		str := c.oset.String()
		if str != c.expected {
			t.Errorf("Expected %v. Got %v", c.expected, str)
		}
	}

}

func TestSlice(t *testing.T) {
	cases := []struct {
		slice []interface{}
	}{
		{[]interface{}{1, 5, -76}},
	}
	for _, c := range cases {
		arr, arrSync := New(c.slice...), NewSync(c.slice...)
		s := arr.Slice()
		for i := range s {
			if s[i] != c.slice[i] {
				t.Errorf("Expected %v. Got %v.", c.slice, s)
			}
		}
		s = arrSync.Slice()
		for i := range s {
			if s[i] != c.slice[i] {
				t.Errorf("Expected %v. Got %v.", c.slice, s)
			}
		}
	}
}

func TestCopyOset(t *testing.T) {
	cases := []struct {
		oset Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.oset.CopyOset()
		if !cpy.IsEqual(c.oset) {
			t.Errorf("Expected %v. Got %v.", c.oset.Slice(), cpy.Slice())
		}
	}
}

func TestCopyArr(t *testing.T) {
	cases := []struct {
		oset Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.oset.CopyArr()
		if !cpy.IsEqual(c.oset) {
			t.Errorf("Expected %v. Got %v.", c.oset.Slice(), cpy.Slice())
		}
	}
}

func TestCopySet(t *testing.T) {
	cases := []struct {
		oset Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.oset.CopySet()
		if !cpy.IsEqual(c.oset) {
			t.Errorf("Expected %v. Got %v.", c.oset.Slice(), cpy.Slice())
		}
	}
}

func TestCopyCollection(t *testing.T) {
	cases := []struct {
		oset Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.oset.CopyCollection()
		if !cpy.IsEqual(c.oset) {
			t.Errorf("Expected %v. Got %v.", c.oset.Slice(), cpy.Slice())
		}
	}
}

func TestArr(t *testing.T) {
	cases := []struct {
		s Interface
		a array.Interface
	}{
		{New(1, 42, -8), array.New(1, 42, -8)},
		{NewSync(1, 42, -8), array.NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		a := c.s.Arr()
		if !a.IsEqual(c.a) {
			t.Errorf("Expected %v. Got %v.", c.a, a)
		}
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		os Interface
		s  set.Interface
	}{
		{New(1, 42, -8), set.New(1, 42, -8)},
		{NewSync(1, 42, -8), set.NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		s := c.os.Set()
		if !s.IsEqual(c.s) {
			t.Errorf("Expected %v. Got %v.", c.s, s)
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		osets    []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(5, 42, 6), New(1, 42, -8, 7)}, New(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{New(1, 42, -8), New(5, 42, 6)}, New(1, 42, -8, 5, 6)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(5, 42, 6), NewSync(1, 42, -8, 7)}, NewSync(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(5, 42, 6)}, NewSync(1, 42, -8, 5, 6)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Union(c.osets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		osets    []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 6, 6), New(1, 7)}, New(42)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(42)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 6, 6), NewSync(1, 7)}, NewSync(42)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(42)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Difference(c.osets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		osets    []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6), New(1, 7)}, New(1)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(1, -8)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6), NewSync(1, 7)}, NewSync(1)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(1, -8)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Intersection(c.osets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestExclusion(t *testing.T) {
	cases := []struct {
		osets    []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6), New(1, 7)}, New(42, 6, 7)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(42, 6)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6), NewSync(1, 7)}, NewSync(42, 6, 7)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(42, 6)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Exclusion(c.osets...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestSort(t *testing.T) {
	less := func(slice []interface{}, i, j int) bool {
		return slice[i].(int) < slice[j].(int)
	}
	cases := []struct {
		oset, sorted Sorted
	}{
		{NewSorted(less, 1, 42, -8), NewSorted(less, -8, 1, 42)},
		{NewSortedSync(less, 1, 42, -8), NewSortedSync(less, -8, 1, 42)},
	}
	for _, c := range cases {
		c.oset.Sort()
		if !c.oset.IsEqual(c.sorted) {
			t.Errorf("Expected %v. Got %v", c.sorted.String(), c.oset.String())
		}
	}
}
