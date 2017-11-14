package main

import "fmt"
import "strconv"
import "github.com/pborman/getopt/v2"
import "github.com/eventi/rpnc/machine"

//var program string = "[d2 0[+oo/o*rso=(s1 0:s2+d4=(1-)ood*>0s)](s)0*+oo=(0*:d./1)]." // find all factors
//var program = "[d2%(3*1+:2/)d1>]"         // collatz conjecture
//var program = "oo/+2/" // approximate sqrt

var debug bool = false
var verbose bool = false
var program = "1^@so!1+^!0@."
var input = ""

func init() {
	getopt.FlagLong(&verbose, "verbose", 'v', "Show verbose output")
	getopt.FlagLong(&debug, "debug", 'd', "Show debug output")
	getopt.FlagLong(&program, "execute", 'e', "Initialize a program")
	getopt.FlagLong(&input, "input", 'i', "Initialize input buffer")
	getopt.Parse()
}

const BYTE = 0
const NUMB = 2
const SKIP = 3

func main() {
	if verbose {
		fmt.Printf("Input: %s\nProgram: %s\n", input, program)
	}
	var m = machine.New()
	m.SendInput(input)
	for _, str := range getopt.Args() {
		val, err := strconv.Atoi(str)
		if err == nil {
			m.Push(val)
		} else {
			panic(err)
		}
	}
	for ix := 0; ix < len(program); ix += 1 {
		bytecode := byte(program[ix])
		if debug {
			m.Debug(program, ix)
		}
		switch m.Mode {
		case NUMB:
			switch {
			case bytecode >= '0' && bytecode <= '9':
				acc := m.Pop()
				acc = acc*10 + int(bytecode-'0')
				m.Push(acc)
			default:
				m.Mode = BYTE
				ix -= 1
			}
		case SKIP: //skip the wrong part of a conditional
			switch bytecode {
			case '(': // nesting - oh no
				m.R.Push(m.Mode)
			case ':':
				m.Mode = m.R.Peek()
			case ')': // end if
				m.Mode = m.R.Pop()
			}
		default:
			switch bytecode {
			// stack manipulation
			case 'd':
				m.Dup()
			case 'o':
				m.Ovr()
			case 'r':
				m.Rol()
			case 's':
				m.Swp()
			case '.':
				m.Dot()
				// math
			case '+':
				m.Add()
			case '-':
				m.Sub()
			case '*':
				m.Mul()
			case '/':
				m.Div()
			case '%':
				m.Mod()
				// comparison
			case '<':
				m.Lst()
			case '=':
				m.Equ()
			case '>':
				m.Gtt()
				// logical negation
			case '~':
				m.Not()
				// conditional
			case '(': // if
				m.R.Push(m.Mode)
				if m.Pop() == 0 {
					m.Mode = SKIP
				}
			case ':': // else
				m.Mode = SKIP
			case ')': // then (end if)
				m.Mode = m.R.Pop()
				// looping
			case '[': // loop
				m.R.Push(ix)
			case ']': // loop
				var condition int
				condition = m.Pop()
				if condition == 0 {
					_ = m.R.Pop()
				} else {
					ix = m.R.Peek()
				}
				// manipulate memory
			case '@':
				m.Fetch()
			case '!':
				m.Store()
			case '#':
				m.Inp()

			case ' ': //noop

			default:
				switch {
				case bytecode >= '0' && bytecode <= '9':
					m.Push(int(bytecode - '0'))
					m.Mode = NUMB
				default:
					panic("Unknown bytecode: [" + string(bytecode) + "]")
				}
			}
		}
	}
	if verbose {
		fmt.Printf("Machine: %s\n", m)
	}
}
