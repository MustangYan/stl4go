package stl4go

import (
	"testing"
)

func Test_Queue_Interface(t *testing.T) {
	_ = Container(NewQueue[int]())
}

func Test_Queue_New(t *testing.T) {
	q := NewQueue[int]()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func Test_Queue_Clear(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.Clear()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func Test_Queue_String(t *testing.T) {
	q := NewQueue[int]()
	expectEq(t, q.String(), "Queue[int]")
}

func Test_Queue_PushFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func Test_Queue_PushBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func Test_Queue_PopFront(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.PopFront()
	expectFalse(t, ok)
}
func Test_Queue_PopBack(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.PopBack()
	expectFalse(t, ok)
}

func Test_Queue_PushFront_PopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	v, ok := q.PopFront()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_Queue_PushFront_PopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	v, ok := q.PopBack()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_Queue_PushBack_PopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	v, ok := q.PopFront()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_Queue_PushBack_PopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	v, ok := q.PopBack()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}
