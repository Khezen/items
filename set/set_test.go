package set

import (
	"github.com/khezen/struct/collection"
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3")
	r := NewTS()
	r.Add("3", "4", "5")
	x := New()
	x.Add("5", "6", "7")

	u := collection.Union(s, r, x)
	if u.Len() != 7 {
		t.Error("Union: the merged set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "3", "4", "5", "6", "7") {
		t.Error("Union: merged items are not availabile in the set.")
	}

	z := collection.Union(x, r)
	if z.Len() != 5 {
		t.Error("Union: Union of 2 sets doesn't have the proper number of items.")
	}
	if settype := reflect.TypeOf(z).String(); settype != "*set.set" {
		t.Error("Union should derive its set type from the first passed set, got", settype)
	}

}

func TestDifference(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3")
	r := NewTS()
	r.Add("3", "4", "5")
	x := New()
	x.Add("5", "6", "7")

	u := collection.Difference(s, r, x)

	if u.Len() != 2 {
		t.Error("Difference: the set doesn't have all items in it.")
	}

	if !u.Has("1", "2") {
		t.Error("Difference: items are not availabile in the set.")
	}

	y := collection.Difference(r, r)
	if y.Len() != 0 {
		t.Error("Difference: size should be zero")
	}

}

func TestIntersection(t *testing.T) {
	s1 := NewTS()
	s1.Add("1", "3", "4", "5")
	s2 := NewTS()
	s2.Add("3", "5", "6")
	s3 := NewTS()
	s3.Add("4", "5", "6", "7")
	u := collection.Intersection(s1, s2, s3)

	if u.Len() != 1 {
		t.Error("Intersection: the set doesn't have all items in it.")
	}

	if !u.Has("5") {
		t.Error("Intersection: items after intersection are not availabile in the set.")
	}
}

func TestIntersection2(t *testing.T) {
	s1 := NewTS()
	s1.Add("1", "3", "4", "5")
	s2 := NewTS()
	s2.Add("5", "6")
	i := collection.Intersection(s1, s2)

	if i.Len() != 1 {
		t.Error("Intersection: size should be 1, it was", i.Len())
	}

	if !i.Has("5") {
		t.Error("Intersection: items after intersection are not availabile in the set.")
	}
}

func TestExclusion(t *testing.T) {
	s := NewTS()
	s.Add("1", "2", "3")
	r := NewTS()
	r.Add("3", "4", "5")
	u := collection.Exclusion(s, r)

	if u.Len() != 4 {
		t.Error("Exclusion: the set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "4", "5") {
		t.Error("Exclusion: items are not availabile in the set.")
	}
}

func BenchmarkSetEquality(b *testing.B) {
	s := NewTS()
	u := NewTS()

	for i := 0; i < b.N; i++ {
		s.Add(i)
		u.Add(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.IsEqual(u)
	}
}

func BenchmarkSubset(b *testing.B) {
	s := NewTS()
	u := NewTS()

	for i := 0; i < b.N; i++ {
		s.Add(i)
		u.Add(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.IsSubset(u)
	}
}

func benchmarkIntersection(b *testing.B, numberOfItems int) {
	s1 := NewTS()
	s2 := NewTS()

	for i := 0; i < numberOfItems/2; i++ {
		s1.Add(i)
	}
	for i := 0; i < numberOfItems; i++ {
		s2.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collection.Intersection(s1, s2)
	}
}

func TestReplace(t *testing.T) {
	cases := []struct {
		set, expected    Interface
		item, substitute interface{}
	}{
		{New(1, 42, -8), New(1, 160, -8), 42, 160},
		{NewTS(1, 42, -8), NewTS(1, 160, -8), 42, 160},
	}
	for _, c := range cases {
		c.set.Replace(c.item, c.substitute)
		if !c.set.IsEqual(c.expected) {
			t.Errorf("Expected %v. Got %v.", c.expected.Slice(), c.set.Slice())
		}
	}
}

func BenchmarkIntersection10(b *testing.B) {
	benchmarkIntersection(b, 10)
}

func BenchmarkIntersection100(b *testing.B) {
	benchmarkIntersection(b, 100)
}

func BenchmarkIntersection1000(b *testing.B) {
	benchmarkIntersection(b, 1000)
}

func BenchmarkIntersection10000(b *testing.B) {
	benchmarkIntersection(b, 10000)
}

func BenchmarkIntersection100000(b *testing.B) {
	benchmarkIntersection(b, 100000)
}

func BenchmarkIntersection1000000(b *testing.B) {
	benchmarkIntersection(b, 1000000)
}
