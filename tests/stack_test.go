package tests

import (
	"testing"

	"github.com/mikeTwoTimes/dsa/lists"
)

func TestStackInt(t *testing.T) {
	s := lists.NewStack[int]()

	// Empty checks
	if s.Len() != 0 {
		t.Fatalf("expected empty stack length 0, got %d", s.Len())
	}

	// Push
	for i := 0; i < 5; i++ {
		s.Push(i)
		if s.Len() != i+1 {
			t.Fatalf("expected len %d, got %d", i+1, s.Len())
		}
	}

	// Pop using peek to verify LIFO
	for i := 4; i >= 0; i-- {
		if s.Peek() != i {
			t.Fatalf("peek expected %d, got %d", i, s.Peek())
		}
		s.Pop()
	}

	if s.Len() != 0 {
		t.Fatalf("stack should be empty")
	}

	// Interleave
	s.Push(10)
	s.Push(20)

	if s.Peek() != 20 {
		t.Fatalf("expected 20")
	}
	s.Pop()

	s.Push(30)

	if s.Peek() != 30 {
		t.Fatalf("expected 30")
	}
	s.Pop()

	if s.Peek() != 10 {
		t.Fatalf("expected 10")
	}
	s.Pop()
}
