package main

import (
	"math/rand"
	"sync"
	"testing"
)

func TestEnqueueBlocking(t *testing.T) {
	q := &Queue{}
	Initialize(q)

	Enqueue(q, 10)

	if q.Head.Next.Value != 10 {
		t.Errorf("Enqueue failed to add value to the queue")
	}

	Enqueue(q, 20)

	if q.Tail.Value != 20 {
		t.Errorf("Enqueue failed to add value to the tail of the queue")
	}
}

func TestDequeueBlocking(t *testing.T) {
	q := &Queue{}
	Initialize(q)

	Enqueue(q, 10)

	if !Dequeue(q) {
		t.Errorf("Dequeue failed to remove element")
	}

	if q.Head.Next != nil || q.Tail.Next != nil {
		t.Errorf("Queue not empty after dequeue")
	}

	Enqueue(q, 10)
	Enqueue(q, 20)

	if !Dequeue(q) {
		t.Errorf("Dequeue failed to remove element")
	}

	if q.Head.Next.Value != 20 {
		t.Errorf("Dequeue removed wrong element")
	}
	Dequeue(q)
	if Dequeue(q) {
		t.Errorf("Dequeue returned true for empty queue")
	}
}

func TestEnqueueDequeueNonBlocking(t *testing.T) {
	q := &Queue_t{}
	Initialize_t(q)

	value := 10
	Enqueue_t(q, value)

	dequeuedValue := Dequeue_t(q)
	if !dequeuedValue {
		t.Error("Dequeue didn't take place")
	}
}

func TestConcurrentEnqueueDequeueNonBlocking(t *testing.T) {
	q := &Queue_t{}
	Initialize_t(q)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			Enqueue_t(q, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			Dequeue_t(q)
		}
	}()

	wg.Wait()

	if q.Head == q.Tail {
		t.Errorf("Queue is not empty after concurrent operations")
	}
}

func BenchmarkNonBlockingQueueAlgorithm(b *testing.B) {
	q := &Queue_t{}
	Initialize_t(q)

	var wg sync.WaitGroup

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				value := rand.Intn(100)
				Enqueue_t(q, value)
				Dequeue_t(q)
			}()
		}
	})

	wg.Wait()
}

func BenchmarkBlockingQueueAlgorithm(b *testing.B) {
	q := &Queue{}
	Initialize(q)

	var wg sync.WaitGroup

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				value := rand.Intn(100)
				Enqueue(q, value)
				Dequeue(q)
			}()
		}
	})

	wg.Wait()
}
