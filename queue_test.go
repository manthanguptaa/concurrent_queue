package main

import (
	"math/rand"
	"sync"
	"testing"
)

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
