package machine

import "testing"
import "fmt"

var tm *Machine

func init() {
	tm = New()
}

func testop(t *testing.T, tm *Machine, ints ...int) bool {
	for _, want := range ints {
		got := tm.Pop()
		if want != got {
			t.Error(fmt.Sprintf("got %v, want %v\n", got, want))
			return false
		}
	}
	return true
}

func Test_dup(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Dup()
	testop(t, tm, 3, 3)
}

func Test_ovr(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Ovr()
	testop(t, tm, 2, 3, 2)
}

func Test_swp(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Swp()
	testop(t, tm, 2, 3)
}

func Test_rol(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Rol()
	testop(t, tm, 1, 3, 2)
}

func Test_drp(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Drp()
	testop(t, tm, 2)
}
func Test_div(t *testing.T) {
	tm := New()
	tm.Push(99)
	tm.Push(33)
	tm.Div()
	if tm.Pop() != 3 {
		t.Error("div didn't work")
	}
}

func Test_add(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Add()
	testop(t, tm, 5)
}

func Test_sub(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Sub()
	testop(t, tm, -1)
}

func Test_mul(t *testing.T) {
	tm := New()
	tm.Push(0, 1, 2, 3)
	tm.Mul()
	testop(t, tm, 6)
}

/*
 *func mod(stack *stack.Stack) {
 *  var a, b int
 *  b = stack.Pop()
 *  a = stack.Pop()
 *  stack.Push(a % b)
 *}
 *
 *func lst(stack *stack.Stack) {
 *  var a, b int
 *  b = stack.Pop()
 *  a = stack.Pop()
 *  if a < b {
 *    stack.Push(1)
 *  } else {
 *    stack.Push(0)
 *  }
 *}
 *
 *func gtt(stack *stack.Stack) {
 *  var a, b int
 *  b = stack.Pop()
 *  a = stack.Pop()
 *  if a > b {
 *    stack.Push(1)
 *  } else {
 *    stack.Push(0)
 *  }
 *}
 *
 *func not(stack *stack.Stack) {
 *  var a int
 *  a = stack.Pop()
 *  if a == 0 {
 *    stack.Push(1)
 *  } else {
 *    stack.Push(0)
 *  }
 *}
 *
 *func equ(stack *stack.Stack) {
 *  var a, b int
 *  b = stack.Pop()
 *  a = stack.Pop()
 *  if a == b {
 *    stack.Push(1)
 *  } else {
 *    stack.Push(0)
 *  }
 *}
 *
 *func fet(stack *stack.Stack) {
 *  // heap contains chars and we need an int
 *  addr := stack.Pop()
 *  //get 4 bytes (I think)
 *  value := 0
 *  for i := 0; i <= 3; i++ {
 *    value = value << 8
 *    value += int(heap[addr+i])
 *  }
 *}
 *
 *func sto(stack *stack.Stack) {
 *  addr := stack.Pop()
 *  fmt.Printf("addr:%v heap:%s\n", addr, heap)
 *  if addr+3 >= len(heap) {
 *    heap = append(heap, make([]byte, 1024)...)
 *  }
 *  val := stack.Pop()
 *  for i := uint(0); i <= 3; i++ {
 *    heap[addr] = byte(val >> i * 8 & 0x000000FF)
 *  }
 *}
 */
