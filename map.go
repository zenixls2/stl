// +build cgo

package stl

// #cgo darwin LDFLAGS: -lstdc++ -L ./
// #cgo darwin CFLAGS: -O2
// #include <stdint.h>
// #include "map.h"
import "C"
import (
	"reflect"
)

type Map struct {
	m C.Map
}

// creates a new empty map
func NewMap() Map {
	var ret Map
	ret.m = C.MapInit()
	return ret
}

// Free release memory hold by map containers.
// be sure to call this to prevent memory leak.
func (sm Map) Free() {
	C.MapFree(sm.m)
}

// Add key, value pair to map
func (sm Map) Add(key int, value interface{}) {
	C.MapAdd(sm.m, C.int(key), C.uintptr_t(reflect.ValueOf(value).Pointer()))
}

// Get value by key from map
func (sm Map) Get(key int) uintptr {
	return uintptr(C.MapGet(sm.m, C.int(key)))
}

// returns an iterator to the first element not less than the given key.
func (sm Map) LowerBound(key int) Iterator {
	iter := MapIterator{}
	iter.i = C.MapLowerBound(sm.m, C.int(key))
	return iter
}

// returns an iterator to the first element greater than the given key
func (sm Map) UpperBound(key int) Iterator {
	iter := MapIterator{}
	iter.i = C.MapLowerBound(sm.m, C.int(key))
	return iter
}
