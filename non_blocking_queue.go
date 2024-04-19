package main

import (
	"sync/atomic"
	"unsafe"
)

type Node_t struct {
	Value int
	Next  *Node_t
}

type Queue_t struct {
	Head *Node_t
	Tail *Node_t
}

func Initialize_t(Q *Queue_t) {
	node := &Node_t{Next: nil}
	Q.Head = node
	Q.Tail = node
}

func Enqueue_t(Q *Queue_t, value int) {
	node := &Node_t{Value: value}
	node.Next = nil
	for {
		tail := Q.Tail
		next := tail.Next
		if tail == Q.Tail {
			if next == nil {
				if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.Next)),
					unsafe.Pointer(next),
					unsafe.Pointer(node)) {
					break
				}
			} else {
				atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&Q.Tail)),
					unsafe.Pointer(tail),
					unsafe.Pointer(next))
			}
		}
	}
}

func Dequeue_t(Q *Queue_t) bool {
	for {
		head := Q.Head
		tail := Q.Tail
		next := head.Next
		if head == Q.Head {
			if head == tail {
				if next == nil {
					return false
				}
				atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&Q.Tail)),
					unsafe.Pointer(tail), unsafe.Pointer(next))
			} else {
				if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&Q.Head)),
					unsafe.Pointer(head), unsafe.Pointer(next)) {
					break
				}
			}
		}
	}
	return true
}
