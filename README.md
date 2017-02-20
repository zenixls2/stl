# stl
--
    import "github.com/zenixls2/stl"

+build cgo

    The stl package provides c++ standard template library binding to golang

Currently this project only demonstrate several techniques to design the
wrapper. Binding to C++ techniques are based on
http://stackoverflow.com/questions/1713214/how-to-use-c-in-go .

### Full Documentation

    https://godoc.org/github.com/zenixls2/stl

### Doc Generation

    $ godocdown > README.md

### Installation

    $ go get github.com/zenixls2/stl

## Usage

#### type Iterator

```go
type Iterator interface {
	Value() uintptr
	Key() int
	Next() Iterator
	Prev() Iterator
}
```


#### type Map

```go
type Map struct {
}
```


#### func  NewMap

```go
func NewMap() Map
```
creates a new empty map

#### func (Map) Add

```go
func (sm Map) Add(key int, value interface{})
```
Add key, value pair to map

#### func (Map) Free

```go
func (sm Map) Free()
```
Free release memory hold by map containers. be sure to call this to prevent
memory leak.

#### func (Map) Get

```go
func (sm Map) Get(key int) uintptr
```
Get value by key from map

#### func (Map) LowerBound

```go
func (sm Map) LowerBound(key int) Iterator
```
returns an iterator to the first element not less than the given key.

#### func (Map) UpperBound

```go
func (sm Map) UpperBound(key int) Iterator
```
returns an iterator to the first element greater than the given key

#### type MapIterator

```go
type MapIterator struct {
}
```


#### func (MapIterator) Key

```go
func (si MapIterator) Key() int
```
get element key from iterator

#### func (MapIterator) Next

```go
func (si MapIterator) Next() Iterator
```
iterator move to next

#### func (MapIterator) Prev

```go
func (si MapIterator) Prev() Iterator
```
iterator move to prev

#### func (MapIterator) Value

```go
func (si MapIterator) Value() uintptr
```
get element pointer from iterator
