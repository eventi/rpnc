package stack

import "testing"

func TestPushPop(t *testing.T) {
	stack := New()
	stack.Push(1)
	tos := stack.Pop()
	if tos != 1 {
		t.Error("Expected 1, got ", tos)
	}
}

func TestPushPeek(t *testing.T) {
	stack := New()
	stack.Push(1)
	tos := stack.Peek()
	if tos != 1 {
		t.Error("Expected 1, got ", tos)
	}
}

func TestPushLen(t *testing.T) {
	stack := New()
	stack.Push(1)
	len := stack.Len()
	if len != 1 {
		t.Error("Expected 1, got ", len)
	}
}
