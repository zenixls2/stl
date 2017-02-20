package stl

type Iterator interface {
	Value() uintptr
	Key() int
	Next() Iterator
	Prev() Iterator
}
