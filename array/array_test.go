package array

import (
	"testing"

	"github.com/khezen/struct/collection"
)

func testErr(err error, expectErr bool, t *testing.T) {
	if (expectErr && err == nil) || (!expectErr && err != nil) {
		t.Errorf(" Error expected? %v. Got: %v.", expectErr, err)
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		array    Interface
		i        int
		expected interface{}
	}{
		{New(1, 7, -5), 2, -5},
		{NewSync(1, 7, -5), 2, -5},
	}
	for _, c := range cases {
		item := c.array.Get(c.i)
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
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8, 42, -1)},
		{NewSync(), New(42, -1), NewSync(42, -1)},
		{NewSync(), New(), NewSync()},
		{NewSync(), New(nil), NewSync(nil)},
		{NewSync(), New(42, -1), NewSync(42, -1)},
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
	}{
		{New(1, 4, -8), New(42, -1), New(1, 4, 42, -1, -8), 2},
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, 42, -1, -8), 2},
	}
	for _, c := range cases {
		c.array.Insert(c.i, c.toBeInserted.Slice()...)
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
		{NewSync(1, 4, -8), New(42, -1), NewSync(1, 4, -8)},
		{NewSync(1, 4, -8), New(1, -8), NewSync(4)},
		{NewSync(), New(42, -1), NewSync()},
		{NewSync(), New(), NewSync()},
		{NewSync(), New(nil), NewSync()},
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
	}{
		{New(1, 4, -8), New(4, -8), 0, 1},
		{NewSync(1, 4, -8), New(4, -8), 0, 1},
	}
	for _, c := range cases {
		removed := c.array.RemoveAt(c.i)
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
		{NewSync(1, 4, -8), NewSync(42, 4, -8), 1, 42},
		{NewSync(1, 4, -8), NewSync(1, 4, -8), 1000, 42},
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
		{NewSync(1, 4, -8), NewSync(1, 42, -8), 1, 42},
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
		{NewSync(1, 42, -8), 42, 1, false},
		{NewSync(1, 42, -8), 1000, -1, true},
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
	}{
		{New(1, 42, -8, 12), New(42, -8), 1, 2},
		{NewSync(1, 42, -8, 12), NewSync(42, -8), 1, 2},
	}
	for _, c := range cases {
		arr := c.array.SubArray(c.i, c.j)
		if !arr.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), arr.Slice())
		}
		if arr.IsEqual(c.array) {
			t.Errorf("c.array should not be modified")
		}
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		array, expected Interface
		i, j            int
	}{
		{New(1, 42, -8), New(42, 1, -8), 0, 1},
		{New(1, 42, -8), New(42, 1, -8), 1, 0},
		{NewSync(1, 42, -8), NewSync(42, 1, -8), 0, 1},
		{NewSync(1, 42, -8), NewSync(42, 1, -8), 1, 0},
	}
	for _, c := range cases {
		c.array.Swap(c.i, c.j)
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
		{NewSync(1, 42, -8), NewSync(1, 42, -8), true},
		{NewSync(1, 42, -8), NewSync(-8), true},
		{NewSync(1, 42, -8), NewSync(34), false},
		{NewSync(1, 42, -8), NewSync(nil), false},
		{NewSync(1, 42, -8), NewSync(), true},
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
		{NewSync(1, -8, 42), 0, 2},
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
		{NewSync(), 0},
		{NewSync(1), 1},
		{NewSync(1, 42, -8), 3},
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
		{NewSync(1, 42, -8)},
		{NewSync()},
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
		{NewSync(), true},
		{NewSync(1, 42, -8), false},
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
		{NewSync(), NewSync(), true},
		{NewSync(1, 42, -8), NewSync(1, 42, -8), true},
		{NewSync(1, 42, -8), NewSync(1, "42", -8), false},
		{NewSync(1, 42, -8), NewSync(), false},
		{NewSync(1, 42, -8), NewSync(42, 1, -8), false},
		{NewSync(66, -1000), NewSync(42, 1, 8), false},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 1, 42, -8), true},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 1, "42", -8), false},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil), false},
		{NewSortedSync(nil, 1, 42, -8), NewSortedSync(nil, 42, 1, -8), false},
		{NewSortedSync(nil, 66, -1000), NewSortedSync(nil, 42, 1, 8), false},
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
		{NewSync(1, 2, 3), "[1 2 3]"},
		{NewSync(-12, 6, 111), "[-12 6 111]"},
		{NewSync(), "[]"},
		{NewSync(nil), "[<nil>]"},
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

func TestCopy(t *testing.T) {
	cases := []struct {
		array Interface
	}{
		{New(1, 42, -8)},
		{New(-66, 1000, 32)},
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
	}
	for _, c := range cases {
		cpy := c.array.CopyArr()
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
		{NewSync(1, 42, -8)},
		{NewSync(-66, 1000, 32)},
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
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(5, 42, 6), NewSync(1, 42, -8, 7)}, NewSync(1, 42, -8, 5, 6, 7)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(5, 42, 6)}, NewSync(1, 42, -8, 5, 6)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
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
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 6, 6), NewSync(1, 7)}, NewSync(42)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(42)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
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
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6), NewSync(1, 7)}, NewSync(1)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(1, -8)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
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
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6), NewSync(1, 7)}, NewSync(42, 6, 7)},
		{[]collection.Interface{NewSync(1, 42, -8), NewSync(-8, 1, 6)}, NewSync(42, 6)},
		{[]collection.Interface{NewSync(1, 42, -8)}, NewSync(1, 42, -8)},
	}
	for _, c := range cases {
		result := collection.Exclusion(c.arrays...)
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
		array, sorted Sorted
	}{
		{NewSorted(less, 1, 42, -8), NewSorted(less, -8, 1, 42)},
		{NewSortedSync(less, 1, 42, -8), NewSortedSync(less, -8, 1, 42)},
	}
	for _, c := range cases {
		c.array.Sort()
		if !c.array.IsEqual(c.sorted) {
			t.Errorf("Expected %v. Got %v", c.sorted.String(), c.array.String())
		}
	}
}
