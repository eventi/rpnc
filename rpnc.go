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

func main() {
	if verbose {
		fmt.Printf("Input: %s\nProgram: %s\n", input, program)
	}
	var m = machine.New()
	m.AcceptInput(input)
	m.SetProgram(program)
	for _, str := range getopt.Args() {
		val, err := strconv.Atoi(str)
		if err == nil {
			m.Push(val)
		} else {
			panic(err)
		}
	}
	end := len(program) + m.IP
	for m.IP != end {
		bytecode := byte(m.H[m.IP])
		if debug {
			m.Debug(string(m.H), m.IP)
		}
		switch m.Mode {
		case machine.NUMB:
			switch {
			case bytecode >= '0' && bytecode <= '9':
				acc := m.Pop()
				acc = acc*10 + int(bytecode-'0')
				m.Push(acc)
			default:
				m.Mode = machine.INTERPRET
				m.IP -= 1
			}
		case machine.SKIP: //skip the wrong part of a conditional
			switch bytecode {
			case '(': // nesting - oh no
				m.R.Push(m.Mode)
			case ':':
				m.Mode = m.R.Peek()
			case ')': // end if
				m.Mode = m.R.Pop()
			}
		case machine.CMPL:
			switch bytecode {
			case '}':
				// end compilation
			case '[':
			case ']':
			case '(':
			case ':':
			case ')':
			case '{':
				panic("you can't create a definition in a definition")
			default:
				//compile
			}
		case machine.INTERPRET: //interpret mode
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
					m.Mode = machine.SKIP
				}
			case ':': // else
				m.Mode = machine.SKIP
			case ')': // then (end if)
				m.Mode = m.R.Pop()
				// looping
			case '[': // loop
				m.R.Push(m.IP)
			case ']': // loop
				var condition int
				condition = m.Pop()
				if condition == 0 {
					_ = m.R.Pop()
				} else {
					m.IP = m.R.Peek()
				}
				// manipulate memory
			case '@':
				m.Fetch()
			case '!':
				m.Store()
			case '#':
				m.Inp()
			case '\'':
				m.Push(int(m.H[m.IP+1]))
				m.IP += 1

			case '{':
				//store the current value of here in the four bytes starting with TOS * 4
				ival := m.Pop()
				for i := 0; i <= 3; i++ {
					m.H[m.Here] = byte(ival & 0x000000FF)
					m.Here++
					ival = ival >> 8
				}
				m.Mode = machine.CMPL
			case '$': //: call    ( token -- )
				m.R.Push(m.IP) //  ip> >R  ( ) (R: ip )
				m.IP = m.Pop() //  >ip     ( )
			case ';':
				m.IP = m.R.Pop()

			case ' ': //noop

			default:
				switch {
				case bytecode >= '0' && bytecode <= '9':
					m.Push(int(bytecode - '0'))
					m.Mode = machine.NUMB
				default:
					panic("Unknown bytecode: [" + string(bytecode) + "]")
				}
			}
		}
		m.IP++
	}
	if verbose {
		fmt.Printf("Machine: %s\n", m)
	}
}
