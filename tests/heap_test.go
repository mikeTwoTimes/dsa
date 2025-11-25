package tests

import (
	"testing"

	"github.com/mikeTwoTimes/dsa/trees"
)

func TestHeapInt(t *testing.T) {
	h := trees.NewHeap[int]()

	if h.Len() != 0 {
		t.Fatalf("expected empty heap")
	}

	values := []int{5, 1, 9, 3, 7, 2}
	for _, v := range values {
		h.Insert(v)
	}

	if h.Len() != len(values) {
		t.Fatalf("heap len mismatch")
	}

	// Pop the heap into sorted order
	sorted := []int{}
	for h.Len() > 0 {
		sorted = append(sorted, h.Peek())
		h.Pop()
	}

	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i-1] {
			t.Fatalf("heap pop produced descending order: %v", sorted)
		}
	}

	// Interleave
	h.Insert(10)
	h.Insert(5)

	if h.Peek() != 5 {
		t.Fatalf("expected 5")
	}
	h.Pop()

	h.Insert(1)
	if h.Peek() != 1 {
		t.Fatalf("expected 1")
	}
	h.Pop()

	if h.Peek() != 10 {
		t.Fatalf("expected 10")
	}
	h.Pop()
}
