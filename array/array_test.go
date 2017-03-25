package array

import (
	"github.com/khezen/check"
	"github.com/khezen/struct/collection"
	"testing"
)

func testErr(err error, expectErr bool, t *testing.T) {
	if !check.ErrorExpectation(err, expectErr) {
		t.Errorf(" Error expected? %v. Got: %v.", expectErr, err)
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		array     Interface
		i         int
		expected  interface{}
		expectErr bool
	}{
		{New(1, 7, -5), 2, -5, false},
		{New(1, 7, -5), -1, nil, true},
		{New(1, 7, -5), 1000, nil, true},
		{NewTS(1, 7, -5), 2, -5, false},
		{NewTS(1, 7, -5), -1, nil, true},
		{NewTS(1, 7, -5), 1000, nil, true},
	}
	for _, c := range cases {
		item, err := c.array.Get(c.i)
		testErr(err, c.expectErr, t)
		if item != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, item)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		array, toBeAdded, expected Interface
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, -8, 42, -1)},
		{New(), New(42, -1), New(42, -1)},
		{New(), New(), New()},
		{New(), New(nil), New(nil)},
		{New(), New(42, -1), New(42, -1)},
		{NewTS(1, 4, -8), New(42, -1), NewTS(1, 4, -8, 42, -1)},
		{NewTS(), New(42, -1), NewTS(42, -1)},
		{NewTS(), New(), NewTS()},
		{NewTS(), New(nil), NewTS(nil)},
		{NewTS(), New(42, -1), NewTS(42, -1)},
	}
	for _, c := range cases {
		c.array.Add(c.toBeAdded.Slice()...)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		array, toBeInserted, expected Interface
		i                             int
		expectErr                     bool
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, 42, -1, -8), 2, false},
		{New(1, 4, -8), New(42, -1), New(1, 4, -8), -1, true},
		{New(1, 4, -8), New(42, -1), New(1, 4, -8), 42, true},
		{NewTS(1, 4, -8), New(42, -1), NewTS(1, 4, 42, -1, -8), 2, false},
		{NewTS(1, 4, -8), New(42, -1), NewTS(1, 4, -8), -1, true},
		{NewTS(1, 4, -8), New(42, -1), NewTS(1, 4, -8), 42, true},
	}
	for _, c := range cases {
		err := c.array.Insert(c.i, c.toBeInserted.Slice()...)
		testErr(err, c.expectErr, t)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		array, toBeRemoved, expected Interface
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, -8)},
		{New(1, 4, -8), New(1, -8), New(4)},
		{New(), New(42, -1), New()},
		{New(), New(), New()},
		{New(), New(nil), New()},
		{NewTS(1, 4, -8), New(42, -1), NewTS(1, 4, -8)},
		{NewTS(1, 4, -8), New(1, -8), NewTS(4)},
		{NewTS(), New(42, -1), NewTS()},
		{NewTS(), New(), NewTS()},
		{NewTS(), New(nil), NewTS()},
	}
	for _, c := range cases {
		c.array.Remove(c.toBeRemoved.Slice()...)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestRemoveAt(t *testing.T) {
	cases := []struct {
		array, expected Interface
		i               int
		removed         interface{}
		expectErr       bool
	}{
		{New(1, 4, -8), New(4, -8), 0, 1, false},
		{New(1, 4, -8), New(1, 4, -8), -1, nil, true},
		{New(1, 4, -8), New(1, 4, -8), 42, nil, true},
		{NewTS(1, 4, -8), New(4, -8), 0, 1, false},
		{NewTS(1, 4, -8), New(1, 4, -8), -1, nil, true},
		{NewTS(1, 4, -8), New(1, 4, -8), 42, nil, true},
	}
	for _, c := range cases {
		removed, err := c.array.RemoveAt(c.i)
		testErr(err, c.expectErr, t)
		if removed != c.removed {
			t.Errorf("Expected %v. Got %v.", c.removed, removed)
		}
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestReplace(t *testing.T) {
	cases := []struct {
		array, expected  Interface
		item, substitute interface{}
	}{
		{New(1, 4, -8), New(42, 4, -8), 1, 42},
		{New(1, 4, -8), New(1, 4, -8), 1000, 42},
		{NewTS(1, 4, -8), NewTS(42, 4, -8), 1, 42},
		{NewTS(1, 4, -8), NewTS(1, 4, -8), 1000, 42},
	}
	for _, c := range cases {
		c.array.Replace(c.item, c.substitute)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestReplaceAt(t *testing.T) {
	cases := []struct {
		array, expected Interface
		i               int
		substitute      interface{}
	}{
		{New(1, 4, -8), New(1, 42, -8), 1, 42},
		{New(1, 4, -8), New(1, 4, -8), 1000, 42},
		{NewTS(1, 4, -8), NewTS(1, 42, -8), 1, 42},
		{NewTS(1, 4, -8), NewTS(1, 4, -8), 1000, 42},
	}
	for _, c := range cases {
		c.array.ReplaceAt(c.i, c.substitute)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v", c.expected.Slice(), c.array.Slice())
		}
	}

}

func TestIndexOf(t *testing.T) {
	cases := []struct {
		array     Interface
		item      interface{}
		i         int
		expectErr bool
	}{
		{New(1, 42, -8), 42, 1, false},
		{New(1, 42, -8), 1000, -1, true},
		{NewTS(1, 42, -8), 42, 1, false},
		{NewTS(1, 42, -8), 1000, -1, true},
	}
	for _, c := range cases {
		i, err := c.array.IndexOf(c.item)
		testErr(err, c.expectErr, t)
		if i != c.i {
			t.Errorf("Expected %v. Got %v.", c.i, i)
		}
	}
}

func TestSubArray(t *testing.T) {
	cases := []struct {
		array, expected Interface
		i, j            int
		expectErr       bool
	}{
		{New(1, 42, -8, 12), New(42, -8), 1, 2, false},
		{New(1, 42, -8, 12), nil, -1, 2, true},
		{New(1, 42, -8, 12), nil, 1000, 2, true},
		{New(1, 42, -8, 12), nil, 1, -2, true},
		{New(1, 42, -8, 12), nil, 1, 1000, true},
		{NewTS(1, 42, -8, 12), NewTS(42, -8), 1, 2, false},
		{NewTS(1, 42, -8, 12), nil, -1, 2, true},
		{NewTS(1, 42, -8, 12), nil, 1000, 2, true},
		{NewTS(1, 42, -8, 12), nil, 1, -2, true},
		{NewTS(1, 42, -8, 12), nil, 1, 1000, true},
	}
	for _, c := range cases {
		arr, err := c.array.SubArray(c.i, c.j)
		testErr(err, c.expectErr, t)
		if !c.expectErr {
			if !arr.IsEqual(c.expected) {
				t.Errorf("Expected %v. Got %v.", c.expected.Slice(), arr.Slice())
			}
			if arr.IsEqual(c.array) {
				t.Errorf("c.array should not be modified")
			}
		}
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		array, expected Interface
		i, j            int
		expectErr       bool
	}{
		{New(1, 42, -8), New(42, 1, -8), 0, 1, false},
		{New(1, 42, -8), New(42, 1, -8), 1, 0, false},
		{New(1, 42, -8), New(1, 42, -8), -1, 0, true},
		{New(1, 42, -8), New(1, 42, -8), 1000, 0, true},
		{New(1, 42, -8), New(1, 42, -8), 1, 1000, true},
		{New(1, 42, -8), New(1, 42, -8), 1, -1, true},
		{NewTS(1, 42, -8), NewTS(42, 1, -8), 0, 1, false},
		{NewTS(1, 42, -8), NewTS(42, 1, -8), 1, 0, false},
		{NewTS(1, 42, -8), NewTS(1, 42, -8), -1, 0, true},
		{NewTS(1, 42, -8), NewTS(1, 42, -8), 1000, 0, true},
		{NewTS(1, 42, -8), NewTS(1, 42, -8), 1, 1000, true},
		{NewTS(1, 42, -8), NewTS(1, 42, -8), 1, -1, true},
	}
	for _, c := range cases {
		err := c.array.Swap(c.i, c.j)
		testErr(err, c.expectErr, t)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestHas(t *testing.T) {
	cases := []struct {
		array, items Interface
		expected     bool
	}{
		{New(1, 42, -8), New(1, 42, -8), true},
		{New(1, 42, -8), New(-8), true},
		{New(1, 42, -8), New(34), false},
		{New(1, 42, -8), New(nil), false},
		{New(1, 42, -8), New(), true},
		{NewTS(1, 42, -8), NewTS(1, 42, -8), true},
		{NewTS(1, 42, -8), NewTS(-8), true},
		{NewTS(1, 42, -8), NewTS(34), false},
		{NewTS(1, 42, -8), NewTS(nil), false},
		{NewTS(1, 42, -8), NewTS(), true},
	}
	for _, c := range cases {
		has := c.array.Has(c.items.Slice()...)
		if has != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, has)
		}
	}
}

func TestEach(t *testing.T) {
	cases := []struct {
		array             Interface
		counter, expected int
	}{
		{New(1, -8, 42), 0, 2},
		{NewTS(1, -8, 42), 0, 2},
	}
	for _, c := range cases {
		callback := func(item interface{}) bool {
			value := item.(int)
			c.counter++
			return value > 0
		}
		c.array.Each(callback)
		if c.counter != c.expected {
			t.Errorf("Expected %v. Got %v.", c.expected, c.counter)
		}
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		array Interface
		len   int
	}{
		{New(), 0},
		{New(1), 1},
		{New(1, 42, -8), 3},
		{NewTS(), 0},
		{NewTS(1), 1},
		{NewTS(1, 42, -8), 3},
	}
	for _, c := range cases {
		if c.array.Len() != c.len {
			t.Errorf("Expected %v. Got %v.", c.len, c.array.Len())
		}
	}
}

func TestClear(t *testing.T) {
	cases := []struct {
		array Interface
	}{
		{New(1, 42, -8)},
		{New()},
		{NewTS(1, 42, -8)},
		{NewTS()},
	}
	for _, c := range cases {
		c.array.Clear()
		if !c.array.IsEmpty() {
			t.Error("Array should be empty")
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		array   Interface
		isEmpty bool
	}{
		{New(), true},
		{New(1, 42, -8), false},
		{NewTS(), true},
		{NewTS(1, 42, -8), false},
	}
	for _, c := range cases {
		if c.array.IsEmpty() != c.isEmpty {
			t.Errorf("Expected %v. Got %v.", c.isEmpty, c.array.IsEmpty())
		}
	}
}

func TestIsEqual(t *testing.T) {
	cases := []struct {
		array, toBeCompared Interface
		isEqual             bool
	}{
		{New(), New(), true},
		{New(1, 42, -8), New(1, 42, -8), true},
		{New(1, 42, -8), New(1, "42", -8), false},
		{New(1, 42, -8), New(), false},
		{New(1, 42, -8), New(42, 1, -8), false},
		{New(66, -1000), New(42, 1, 8), false},
		{NewTS(), NewTS(), true},
		{NewTS(1, 42, -8), NewTS(1, 42, -8), true},
		{NewTS(1, 42, -8), NewTS(1, "42", -8), false},
		{NewTS(1, 42, -8), NewTS(), false},
		{NewTS(1, 42, -8), NewTS(42, 1, -8), false},
		{NewTS(66, -1000), NewTS(42, 1, 8), false},
	}
	for _, c := range cases {
		isEqual := c.array.IsEqual(c.toBeCompared)
		if isEqual != c.isEqual {
			t.Errorf("Expected %v to be equal to %v? %v. Got: %v", c.array.Slice(), c.toBeCompared.Slice(), c.isEqual, isEqual)
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		array, toBeMerged, expected Interface
	}{
		{New(1, 42), New(-8), New(1, 42, -8)},
		{New(1, 42), New(-8, nil), New(1, 42, -8, nil)},
		{New(1, 42), New(), New(1, 42)},
		{New(), New(), New()},
		{NewTS(1, 42), NewTS(-8), NewTS(1, 42, -8)},
		{NewTS(1, 42), NewTS(-8, nil), NewTS(1, 42, -8, nil)},
		{NewTS(1, 42), NewTS(), NewTS(1, 42)},
		{NewTS(), NewTS(), NewTS()},
	}
	for _, c := range cases {
		c.array.Merge(c.toBeMerged)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestSeparate(t *testing.T) {
	cases := []struct {
		array, toBeMerged, expected Interface
	}{
		{New(1, 42, -8), New(1, 42), New(-8)},
		{New(1, 42, -8), New(1, 42, nil), New(-8)},
		{New(1, 42, -8), New(), New(1, 42, -8)},
		{New(), New(), New()},
		{NewTS(1, 42, -8), NewTS(1, 42), NewTS(-8)},
		{NewTS(1, 42, -8), NewTS(1, 42, nil), NewTS(-8)},
		{NewTS(1, 42, -8), NewTS(), NewTS(1, 42, -8)},
		{NewTS(), NewTS(), NewTS()},
	}
	for _, c := range cases {
		c.array.Separate(c.toBeMerged)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestRetain(t *testing.T) {
	cases := []struct {
		array, toBeMerged, expected Interface
	}{
		{New(1, 42, -8), New(1, -8, 100), New(1, -8)},
		{New(1, 42, -8), New(1, -8, 100, nil), New(1, -8)},
		{New(1, 42, -8), New(), New()},
		{New(), New(), New()},
		{NewTS(1, 42, -8), NewTS(1, -8, 100), NewTS(1, -8)},
		{NewTS(1, 42, -8), NewTS(1, -8, 100, nil), NewTS(1, -8)},
		{NewTS(1, 42, -8), NewTS(), NewTS()},
		{NewTS(), NewTS(), NewTS()},
	}
	for _, c := range cases {
		c.array.Retain(c.toBeMerged)
		if !c.array.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.array.Slice())
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		array    Interface
		expected string
	}{
		{New(1, 2, 3), "[1 2 3]"},
		{New(-12, 6, 111), "[-12 6 111]"},
		{New(), "[]"},
		{New(nil), "[<nil>]"},
		{NewTS(1, 2, 3), "[1 2 3]"},
		{NewTS(-12, 6, 111), "[-12 6 111]"},
		{NewTS(), "[]"},
		{NewTS(nil), "[<nil>]"},
	}

	for _, c := range cases {
		str := c.array.String()
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
		arr, arrTS := New(c.slice...), NewTS(c.slice...)
		s := arr.Slice()
		for i := range s {
			if s[i] != c.slice[i] {
				t.Errorf("Expected %v. Got %v.", c.slice, s)
			}
		}
		s = arrTS.Slice()
		for i := range s {
			if s[i] != c.slice[i] {
				t.Errorf("Expected %v. Got %v.", c.slice, s)
			}
		}
	}
}

func TestCopy(t *testing.T) {
	cases := []struct {
		array Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewTS(1, 42, -8)},
		{NewTS(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.array.Copy()
		if !cpy.IsEqual(c.array) {
			t.Errorf("Expected %v. Got %v.", c.array.Slice(), cpy.Slice())
		}
	}
}

func TestCopyCollection(t *testing.T) {
	cases := []struct {
		array Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewTS(1, 42, -8)},
		{NewTS(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.array.CopyCollection()
		if !cpy.IsEqual(c.array) {
			t.Errorf("Expected %v. Got %v.", c.array.Slice(), cpy.Slice())
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		arrays   []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(5, 42, 6), New(1, 42, -8, 7)}, New(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{New(1, 42, -8), New(5, 42, 6)}, New(1, 42, -8, 5, 6)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(5, 42, 6), NewTS(1, 42, -8, 7)}, NewTS(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(5, 42, 6)}, NewTS(1, 42, -8, 5, 6)},
		{[]collection.Interface{NewTS(1, 42, -8)}, NewTS(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Union(c.arrays...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		arrays   []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 6, 6), New(1, 7)}, New(42)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(42)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(-8, 6, 6), NewTS(1, 7)}, NewTS(42)},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(-8, 1, 6)}, NewTS(42)},
		{[]collection.Interface{NewTS(1, 42, -8)}, NewTS(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Difference(c.arrays...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		arrays   []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6), New(1, 7)}, New(1)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(1, -8)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(-8, 1, 6), NewTS(1, 7)}, NewTS(1)},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(-8, 1, 6)}, NewTS(1, -8)},
		{[]collection.Interface{NewTS(1, 42, -8)}, NewTS(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Intersection(c.arrays...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}

func TestExclusion(t *testing.T) {
	cases := []struct {
		arrays   []collection.Interface
		expected Interface
	}{
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6), New(1, 7)}, New(42, 6, 7)},
		{[]collection.Interface{New(1, 42, -8), New(-8, 1, 6)}, New(42, 6)},
		{[]collection.Interface{New(1, 42, -8)}, New(1, 42, -8)},
		{[]collection.Interface{}, nil},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(-8, 1, 6), NewTS(1, 7)}, NewTS(42, 6, 7)},
		{[]collection.Interface{NewTS(1, 42, -8), NewTS(-8, 1, 6)}, NewTS(42, 6)},
		{[]collection.Interface{NewTS(1, 42, -8)}, NewTS(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Exclusion(c.arrays...)
		if c.expected != nil && !result.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
		}
	}
}
