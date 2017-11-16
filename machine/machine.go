package machine

import "github.com/eventi/rpnc/stack"
import "fmt"

const INTERPRET = 0
const CMPL = 1
const NUMB = 2
const SKIP = 3

type Machine struct {
	s     stack.Stack
	R     stack.Stack
	H     []byte
	Mode  int
	Here  int
	input []byte
	start int
	IP    int
}

func New() *Machine {
	this := new(Machine)
	this.H = make([]byte, 1024)
	this.Here = 512
	return this
}

func (this *Machine) AcceptInput(input string) {
	for _, b := range input {
		this.input = append(this.input, byte(b))
	}
}

func (this *Machine) SetProgram(input string) {
	this.IP = this.Here
	for _, b := range input {
		this.H[this.Here] = byte(b)
		this.Here++
	}
}

func (this *Machine) Debug(program string, ix int) {
	var modestr string
	switch this.Mode {
	case INTERPRET:
		modestr = "interpret"
	case NUMB:
		modestr = "number"
	case SKIP:
		modestr = "skip"
	}

	fmt.Printf("%s %s %s{%s}%s %d (r: %s)\n", string(modestr[0]), this.s, program[:ix], string(program[ix]), program[ix+1:], ix, this.R)
}

func (this *Machine) Push(ints ...int) {
	for _, i := range ints {
		this.s.Push(i)
	}
}

func (this *Machine) Pop() int {
	return this.s.Pop()
}

func (this *Machine) String() string {
	return fmt.Sprintf("%s", &this.s)
}

//memory functions
func (this *Machine) Fetch() {
	// heap contains chars and we need an int
	addr := this.s.Pop()
	value := 0
	for i := 0; i <= 3; i++ {
		value = value << 8
		value += int(this.H[addr+i])
	}
	this.Push(value)
}

func (this *Machine) Store() {
	addr := this.Pop()
	for addr+3 >= len(this.H) {
		this.H = append(this.H, make([]byte, 1024)...)
	}
	val := this.Pop()
	for i := uint(0); i <= 3; i++ {
		this.H[addr] = byte(val >> i * 8 & 0x000000FF)
	}
}

// The actual bytecodes
func (this *Machine) Dup() {
	this.s.Push(this.s.Peek())
}

func (this *Machine) Drp() {
	this.s.Pop()
}

func (this *Machine) Ovr() {
	var a, b int
	a = this.s.Pop()
	b = this.s.Pop()
	this.s.Push(b)
	this.s.Push(a)
	this.s.Push(b)
}

func (this *Machine) Swp() {
	var a, b int
	a = this.s.Pop()
	b = this.s.Pop()
	this.s.Push(a)
	this.s.Push(b)
}

func (this *Machine) Rol() {
	var a, b, c int
	a = this.s.Pop()
	b = this.s.Pop()
	c = this.s.Pop()
	this.s.Push(b)
	this.s.Push(a)
	this.s.Push(c)
}

func (this *Machine) Dot() {
	fmt.Printf("%d\n", this.s.Pop())
}

func (this *Machine) Add() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	this.s.Push(a + b)
}

func (this *Machine) Sub() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	this.s.Push(a - b)
}

func (this *Machine) Mul() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	this.s.Push(a * b)
}

func (this *Machine) Div() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	this.s.Push(a / b)
}

func (this *Machine) Mod() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	this.s.Push(a % b)
}

func (this *Machine) Lst() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	if a < b {
		this.s.Push(1)
	} else {
		this.s.Push(0)
	}
}

func (this *Machine) Gtt() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	if a > b {
		this.s.Push(1)
	} else {
		this.s.Push(0)
	}
}

func (this *Machine) Not() {
	var a int
	a = this.s.Pop()
	if a == 0 {
		this.s.Push(1)
	} else {
		this.s.Push(0)
	}
}

func (this *Machine) Equ() {
	var a, b int
	b = this.s.Pop()
	a = this.s.Pop()
	if a == b {
		this.s.Push(1)
	} else {
		this.s.Push(0)
	}
}

func (this *Machine) Inp() {
	// get byte from input
	inbyte := this.input[0]
	this.input = this.input[1:]
	this.s.Push(int(inbyte))
}
