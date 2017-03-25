package collection

//
// import (
// 	"github.com/khezen/items/array"
// 	"testing"
// )
//
// func TestUnion(t *testing.T) {
// 	cases := []struct {
// 		arrays   []Interface
// 		expected Interface
// 	}{
// 		{[]Interface{array.New(1, 42, -8), array.New(5, 42, 6), array.New(1, 42, -8, 7)}, array.New(1, 42, -8, 5, 6, 7)},
// 		{[]Interface{array.New(1, 42, -8), array.New(5, 42, 6)}, array.New(1, 42, -8, 5, 6)},
// 		{[]Interface{array.New(1, 42, -8)}, array.New(1, 42, -8)},
// 		{[]Interface{}, nil},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(5, 42, 6), array.NewTS(1, 42, -8, 7)}, array.NewTS(1, 42, -8, 5, 6, 7)},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(5, 42, 6)}, array.NewTS(1, 42, -8, 5, 6)},
// 		{[]Interface{array.NewTS(1, 42, -8)}, array.NewTS(1, 42, -8)},
// 	}
// 	for _, c := range cases {
// 		result := Union(c.arrays...)
// 		if c.expected != nil && !result.IsEqual(c.expected) {
// 			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
// 		}
// 	}
// }
//
// func TestDifference(t *testing.T) {
// 	cases := []struct {
// 		arrays   []Interface
// 		expected Interface
// 	}{
// 		{[]Interface{array.New(1, 42, -8), array.New(-8, 6, 6), array.New(1, 7)}, array.New(42)},
// 		{[]Interface{array.New(1, 42, -8), array.New(-8, 1, 6)}, array.New(42)},
// 		{[]Interface{array.New(1, 42, -8)}, array.New(1, 42, -8)},
// 		{[]Interface{}, nil},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(-8, 6, 6), array.NewTS(1, 7)}, array.NewTS(42)},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(-8, 1, 6)}, array.NewTS(42)},
// 		{[]Interface{array.NewTS(1, 42, -8)}, array.NewTS(1, 42, -8)},
// 	}
// 	for _, c := range cases {
// 		result := Difference(c.arrays...)
// 		if c.expected != nil && !result.IsEqual(c.expected) {
// 			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
// 		}
// 	}
// }
//
// func TestIntersection(t *testing.T) {
// 	cases := []struct {
// 		arrays   []Interface
// 		expected Interface
// 	}{
// 		{[]Interface{array.New(1, 42, -8), array.New(-8, 1, 6), array.New(1, 7)}, array.New(1)},
// 		{[]Interface{array.New(1, 42, -8), array.New(-8, 1, 6)}, array.New(1, -8)},
// 		{[]Interface{array.New(1, 42, -8)}, array.New(1, 42, -8)},
// 		{[]Interface{}, nil},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(-8, 1, 6), array.NewTS(1, 7)}, array.NewTS(1)},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(-8, 1, 6)}, array.NewTS(1, -8)},
// 		{[]Interface{array.NewTS(1, 42, -8)}, array.NewTS(1, 42, -8)},
// 	}
// 	for _, c := range cases {
// 		result := Intersection(c.arrays...)
// 		if c.expected != nil && !result.IsEqual(c.expected) {
// 			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
// 		}
// 	}
// }
//
// func TestExclusion(t *testing.T) {
// 	cases := []struct {
// 		arrays   []Interface
// 		expected Interface
// 	}{
// 		{[]Interface{array.New(1, 42, -8), array.New(-8, 1, 6), array.New(1, 7)}, array.New(42, 6, 7)},
// 		{[]Interface{array.New(1, 42, -8), array.New(-8, 1, 6)}, array.New(42, 6)},
// 		{[]Interface{array.New(1, 42, -8)}, array.New(1, 42, -8)},
// 		{[]Interface{}, nil},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(-8, 1, 6), array.NewTS(1, 7)}, array.NewTS(42, 6, 7)},
// 		{[]Interface{array.NewTS(1, 42, -8), array.NewTS(-8, 1, 6)}, array.NewTS(42, 6)},
// 		{[]Interface{array.NewTS(1, 42, -8)}, array.NewTS(1, 42, -8)},
// 	}
// 	for _, c := range cases {
// 		result := Exclusion(c.arrays...)
// 		if c.expected != nil && !result.IsEqual(c.expected) {
// 			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), result.Slice())
// 		}
// 	}
// }
