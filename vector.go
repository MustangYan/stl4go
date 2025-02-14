package stl4go

// Vector is a sequence container representing array that can change in size.
type Vector[T any] []T

// MakeVector creates an empty Vector object.
func MakeVector[T any]() Vector[T] {
	return (Vector[T])([]T{})
}

// MakeVectorCap creates an empty Vector object with specified capacity.
func MakeVectorCap[T any](c int) Vector[T] {
	v := make([]T, 0, c)
	return (Vector[T])(v)
}

// MakeVectorOf creates an Vector object with initial values.
func MakeVectorOf[T any](v ...T) Vector[T] {
	return (Vector[T])(v)
}

// IsEmpty implements the Container interface.
func (v *Vector[T]) IsEmpty() bool {
	return len(*v) == 0
}

// Len implements the Container interface.
func (v *Vector[T]) Len() int {
	return len(*v)
}

// Cap returns the capacity of the vector.
func (v *Vector[T]) Cap() int {
	return cap(*v)
}

// Clear erases all elements from the vector. After this call, Len() returns zero.
// Leaves the Cap() of the vector unchanged.
func (v *Vector[T]) Clear() {
	*v = (*v)[0:0]
}

// Reserve increases the capacity of the vector (the total number of elements
// that the vector can hold without requiring reallocation)to a value that's
// greater or equal to l. If l is greater than the current Cap(), new storage
// is allocated, otherwise the function does nothing.
//
// Reserve() does not change the size of the vector.
func (v *Vector[T]) Reserve(l int) {
	if cap(*v) < l {
		t := make([]T, len(*v), l)
		copy(t, *v)
		*v = t
	}
}

// Shrink removes unused capacity from the vector.
func (v *Vector[T]) Shrink() {
	if cap(*v) > len(*v) {
		*v = append([]T{}, *v...)
	}
}

// At returns the element value at the index i.
// You can also use the [] operator, and it's better.
func (v *Vector[T]) At(i int) T {
	return (*v)[i]
}

// Set sets the value of the element at the index i.
// You can also use the [] operator, and it's better.
func (v *Vector[T]) Set(i int, x T) {
	(*v)[i] = x
}

// PushBack pushs an element to the end of the vector.
func (v *Vector[T]) PushBack(x T) {
	*v = append(*v, x)
}

// Append appends the values x... to the tail of the vector.
func (v *Vector[T]) Append(x ...T) {
	*v = append(*v, x...)
}

// Insert inserts the values x... into the vector at index i.
// After the insertion, (*v)[i] == x[0].
// Insert panics if i is out of range.
//
// Complexity: O(len(s) + len(v)).
func (v *Vector[T]) Insert(i int, x ...T) {
	s := *v
	total := len(s) + len(x)
	if total <= cap(s) {
		s2 := s[:total]
		copy(s2[i+len(x):], s[i:])
		copy(s2[i:], x)
		*v = s2
		return
	}
	s2 := make([]T, total)
	copy(s2, s[:i])
	copy(s2[i:], x)
	copy(s2[i+len(x):], s[i:])
	*v = s2
}

// Remove removes 1 element in the vector.
//
// Complexity: O(len(s) - i).
func (v *Vector[T]) Remove(i int) {
	v.RemoveRange(i, i+1)
}

// RemoveRange removes the elements in the range[i, j) from the vector.
func (v *Vector[T]) RemoveRange(i, j int) {
	*v = append((*v)[:i], (*v)[j:]...)
}

// RemoveLength removes the elements in the range[i, i+len) from the vector.
func (v *Vector[T]) RemoveLength(i int, len int) {
	v.RemoveRange(i, i+len)
}

// Iterate returns an iterator to the whole vector.
func (v Vector[T]) Iterate() MutableIterator[T] {
	return &vectorIterator[T]{v, 0}
}

// IterateRange returns an iterator to the range [i, j) of the vector.
func (v Vector[T]) IterateRange(i, j int) MutableIterator[T] {
	return &vectorIterator[T]{v[i:j], 0}
}

type vectorIterator[T any] struct {
	v Vector[T]
	i int
}

func (it vectorIterator[T]) Value() T {
	return it.v[it.i]
}

func (it vectorIterator[T]) Pointer() *T {
	return &it.v[it.i]
}

func (it vectorIterator[T]) IsNotEnd() bool {
	return it.i < len(it.v)
}

func (it *vectorIterator[T]) MoveToNext() {
	it.i++
}
