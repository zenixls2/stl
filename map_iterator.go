package stl

// #cgo darwin LDFLAGS: -lstdc++ -L ./
// #cgo darwin CFLAGS: -O2
// #include <stdint.h>
// #include "map.h"
import "C"

type MapIterator struct {
	i C.MapIterator
}

func (si MapIterator) Value() uintptr {
	return uintptr(C.MapIteratorGet(si.i))
}

func (si MapIterator) Key() int {
	return int(C.MapIteratorGetKey(si.i))
}

func (si MapIterator) Next() Iterator {
	if C.MapIteratorNext(si.i) == nil {
		return nil
	}
	return si
}

func (si MapIterator) Prev() Iterator {
	if C.MapIteratorPrev(si.i) == nil {
		return nil
	}
	return si
}
