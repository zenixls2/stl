// +build cgo
package stl

// #cgo darwin LDFLAGS: -lstdc++ -L ./
// #cgo darwin CFLAGS: -O2
// #include <stdint.h>
// #include "map.h"
import "C"

type MapIterator struct {
	i C.MapIterator
}

// get element pointer from iterator
func (si MapIterator) Value() uintptr {
	return uintptr(C.MapIteratorGet(si.i))
}

// get element key from iterator
func (si MapIterator) Key() int {
	return int(C.MapIteratorGetKey(si.i))
}

// iterator move to next
func (si MapIterator) Next() Iterator {
	if C.MapIteratorNext(si.i) == nil {
		return nil
	}
	return si
}

// iterator move to prev
func (si MapIterator) Prev() Iterator {
	if C.MapIteratorPrev(si.i) == nil {
		return nil
	}
	return si
}
