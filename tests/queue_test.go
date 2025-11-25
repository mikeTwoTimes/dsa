package tests

import (
	"testing"

	"github.com/mikeTwoTimes/dsa/lists"
)

func TestQueueInt(t *testing.T) {
	q := lists.NewQueue[int]()

	// Empty queue checks
	if q.Len() != 0 {
		t.Fatalf("expected empty queue length 0, got %d", q.Len())
	}

	// Enqueue
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
		if q.Len() != i+1 {
			t.Fatalf("expected length %d, got %d", i+1, q.Len())
		}
	}

	// Check order by peeking then dequeuing
	for i := 0; i < 5; i++ {
		v := q.Peek()
		if v != i {
			t.Fatalf("peek expected %d, got %d", i, v)
		}
		q.Deque()
	}

	if q.Len() != 0 {
		t.Fatalf("expected empty queue")
	}

	// Interleave
	q.Enqueue(10)
	q.Enqueue(20)

	if q.Peek() != 10 {
		t.Fatalf("expected peek 10")
	}
	q.Deque()

	q.Enqueue(30)

	if q.Peek() != 20 {
		t.Fatalf("expected peek 20")
	}
	q.Deque()

	if q.Peek() != 30 {
		t.Fatalf("expected peek 30")
	}
	q.Deque()
}
