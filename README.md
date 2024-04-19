# Concurrent Queue

This project implements two concurrent queue data structures in Go:

* **Queue with Separate Head and Tail Locks:** This implementation uses a single node structure and separate mutexes for head and tail access to improve concurrency compared to a single lock approach.

* **Non-Blocking Queue:** This implementation uses atomic operations and a lock-free approach for enqueue and dequeue operations using compare and swap. It is suitable for scenarios where high concurrency and avoiding lock overhead is critical.

Both implementations are inspired by the paper "Simple, Fast, and Practical Non-Blocking and Blocking Concurrent Queue Algorithms : [https://www.cs.rochester.edu/~scott/papers/1996_PODC_queues.pdf](https://www.cs.rochester.edu/~scott/papers/1996_PODC_queues.pdf)" by Michael Scott and William Scherer.

**Testing:**

```bash
go test -v ./...
```

**Benchmarking:**

```bash
go test -v -bench=.
```

**Benchmarking Results:**


<img width="662" alt="Screenshot 2024-04-19 at 7 54 01 PM" src="https://github.com/manthanguptaa/concurrent_queue/assets/42516515/79645d94-eb14-4574-bcd0-f3353ee3c0a9">


**Choosing the Right Queue Implementation:**

The choice between these two implementations depends on your specific needs:

* **Queue with Separate Locks:** This offers good performance with balanced concurrency for enqueue and dequeue operations. It's a good choice for general-purpose use cases where thread safety is important.

* **Non-Blocking Queue:** This excels in high-contention scenarios where there's frequent concurrent access. However, it might have a slightly higher overhead compared to the locked version. 

**Additional Notes:**

* The provided tests cover basic enqueue and dequeue operations for both queue implementations.
* Feel free to extend the functionality and tests based on your specific needs.

**Disclaimer:**

These are basic implementations for educational purposes. For production use, consider utilizing more robust and optimized concurrent queue libraries available in the Go ecosystem.
