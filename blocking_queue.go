package main

import "sync"

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	H_lock sync.Mutex
	T_lock sync.Mutex
}

func Initialize(Q *Queue) {
	node := &Node{}
	node.Next = nil
	Q.Head = node
	Q.Tail = node
}

func Enqueue(Q *Queue, value int) {
	node := &Node{Value: value, Next: nil}
	Q.T_lock.Lock()
	defer Q.T_lock.Unlock()
	Q.Tail.Next = node
	Q.Tail = node
}

func Dequeue(Q *Queue) bool {
	Q.H_lock.Lock()
	defer Q.H_lock.Unlock()
	node := Q.Head
	new_head := node.Next
	if new_head == nil {
		return false
	}
	Q.Head = new_head
	return true
}
