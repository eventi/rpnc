package main

import "fmt"
import "strconv"
import "github.com/pborman/getopt/v2"
import "github.com/eventi/rpnc/stack"

var debug bool = false

//var program string = "[d2 0[+oo/o*rso=(s1 0:s2+d4=(1-)ood*>0s)](s)0*+oo=(0*:d./1)]." // find all factors
//var program = "[d2%(3*1+:2/)d1>]"         // collatz conjecture
//var program = "oo/+2/" // approximate sqrt
var program = "1 ^@ so ! 1+ ^! 0@ ."
var heap []int

func init() {
	getopt.Flag(&debug, 'd', "Show debug output")
	getopt.FlagLong(&program, "--execute", 'e', "Initialize a program")
	getopt.Parse()
}

func _debug(mode int, mainstack *stack.Stack, program string, ix int, returnstack *stack.Stack) {
	var modestr string
	switch mode {
	case BYTE:
		modestr = "bytecode"
	case NUMB:
		modestr = "number"
	case COND:
		modestr = "conditional"
	case SKIP:
		modestr = "skip"
	}

	if debug {
		fmt.Printf("%s (%s ) %s{%s}%s %d (r: %s)\n", string(modestr[0]), mainstack, program[:ix], string(program[ix]), program[ix+1:], ix, returnstack)
	}
}

// The actual bytecodes
// each takes a *stack
func dup(stack *stack.Stack) {
	stack.Push(stack.Peek())
}

func ovr(stack *stack.Stack) {
	var a, b int
	a = stack.Pop()
	b = stack.Pop()
	stack.Push(b)
	stack.Push(a)
	stack.Push(b)
}

func swp(stack *stack.Stack) {
	var a, b int
	a = stack.Pop()
	b = stack.Pop()
	stack.Push(a)
	stack.Push(b)
}

func rol(stack *stack.Stack) {
	var a, b, c int
	a = stack.Pop()
	b = stack.Pop()
	c = stack.Pop()
	stack.Push(b)
	stack.Push(a)
	stack.Push(c)
}

func dot(stack *stack.Stack) {
	fmt.Printf("%d\n", stack.Pop())
}

func add(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	stack.Push(a + b)
}

func sub(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	stack.Push(a - b)
}

func mul(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	stack.Push(a * b)
}

func div(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	stack.Push(a / b)
}

func mod(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	stack.Push(a % b)
}

func lst(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	if a < b {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
}

func gtt(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	if a > b {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
}

func not(stack *stack.Stack) {
	var a int
	a = stack.Pop()
	if a == 0 {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
}

func equ(stack *stack.Stack) {
	var a, b int
	b = stack.Pop()
	a = stack.Pop()
	if a == b {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
}

func fet(stack *stack.Stack) {
	stack.Push(heap[stack.Pop()])
}

func sto(stack *stack.Stack) {
	addr := stack.Pop()
	fmt.Printf("addr:%v heap:%s\n", addr, heap)
	switch {
	case addr == len(heap):
		heap = append(heap, stack.Pop())
	case addr > len(heap):
		panic("I just don't know what to do")
	default:
		heap[addr] = stack.Pop()
	}
}

const BYTE = 0
const COND = 1
const NUMB = 2
const SKIP = 3

func main() {
	mainstack := stack.New()
	mode := BYTE
	returnstack := stack.New()
	modestack := stack.New()
	here := 1 // here points to the next available memory cell
	heap = append(heap, here)
	here += 1
	for _, str := range getopt.Args() {
		val, err := strconv.Atoi(str)
		if err == nil {
			mainstack.Push(val)
		} else {
			panic(err)
		}
	}
	for ix := 0; ix < len(program); ix += 1 {
		bytecode := byte(program[ix])
		_debug(mode, mainstack, program, ix, returnstack)
		switch mode {
		case NUMB:
			switch {
			case bytecode >= '0' && bytecode <= '9':
				acc := mainstack.Pop()
				acc = acc*10 + int(bytecode-'0')
				mainstack.Push(acc)
			default:
				mode = BYTE
				ix -= 1
			}
		case SKIP: //skip the wrong part of a conditional
			switch bytecode {
			case '(': // nesting - oh no
				modestack.Push(mode)
			case ':': // else
				mode = COND
			case ')': // end if
				mode = modestack.Pop()
			}
		default:
			switch bytecode {
			// stack manipulation
			case 'd':
				dup(mainstack)
			case 'o':
				ovr(mainstack)
			case 'r':
				rol(mainstack)
			case 's':
				swp(mainstack)
			case '.':
				dot(mainstack)
				// math
			case '+':
				add(mainstack)
			case '-':
				sub(mainstack)
			case '*':
				mul(mainstack)
			case '/':
				div(mainstack)
			case '%':
				mod(mainstack)
				// comparison
			case '<':
				lst(mainstack)
			case '=':
				equ(mainstack)
			case '>':
				gtt(mainstack)
				// logical negation
			case '~':
				not(mainstack)
				// conditional
			case '(': // if
				if mainstack.Pop() == 0 {
					mode = SKIP
				} else {
					mode = COND
				}
			case ':': // else
				mode = SKIP
			case ')': // then (end if)
				mode = modestack.Pop()
			case '[': // loop
				returnstack.Push(ix)
			case ']': // loop
				var condition int
				condition = mainstack.Pop()
				if condition == 0 {
					_ = returnstack.Pop()
				} else {
					ix = returnstack.Peek()
				}
				// manipulate memory
			case '^':
				mainstack.Push(0)
			case '@':
				fet(mainstack)
			case '!':
				sto(mainstack)

			case ' ': //noop

			default:
				switch {
				case bytecode >= '0' && bytecode <= '9':
					mainstack.Push(int(bytecode - '0'))
					mode = NUMB
				default:
					panic("Unknown bytecode: [" + string(bytecode) + "]")
				}
			}
		}
	}
	fmt.Printf("Stack: %s\n", mainstack)
}
