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

func NewMap() Map {
	var ret Map
	ret.m = C.MapInit()
	return ret
}

func (sm Map) Free() {
	C.MapFree(sm.m)
}

func (sm Map) Add(key int, value interface{}) {
	C.MapAdd(sm.m, C.int(key), C.uintptr_t(reflect.ValueOf(value).Pointer()))
}

func (sm Map) Get(key int) uintptr {
	return uintptr(C.MapGet(sm.m, C.int(key)))
}

func (sm Map) LowerBound(key int) Iterator {
	iter := MapIterator{}
	iter.i = C.MapLowerBound(sm.m, C.int(key))
	return iter
}

func (sm Map) UpperBound(key int) Iterator {
	iter := MapIterator{}
	iter.i = C.MapLowerBound(sm.m, C.int(key))
	return iter
}
