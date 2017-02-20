package stl

import "unsafe"
import "testing"

func TestMap(t *testing.T) {
	m := NewMap()
	a := []int{1, 2, 3, 4, 5, 6}
	m.Add(1, &a[0])
	m.Add(3, &a[1])
	m.Add(5, &a[2])
	m.Add(7, &a[3])
	m.Add(9, &a[4])
	m.Add(11, &a[5])
	if *(*int)(unsafe.Pointer(m.Get(1))) != 1 {
		t.Error("Map Get failed")
	}
	iter := m.LowerBound(7)
	index := 3
	for end := iter; end != nil; end = iter.Next() {
		if *(*int)(unsafe.Pointer(iter.Value())) != a[index] {
			t.Error("Map Lower Bound Next() Error")
		}
		index += 1
	}
	iter = m.UpperBound(7)
	index = 3
	for end := iter; end != nil; end = iter.Prev() {
		if *(*int)(unsafe.Pointer(iter.Value())) != a[index] {
			t.Error("Map Lower Bound Next() Error")
		}
		index -= 1
	}
}
