package main

import "fmt"
import "strconv"
import "github.com/pborman/getopt/v2"
import "github.com/eventi/rpnc/stack"


var debug   bool   = false
var program string = "[d2 0[+oo/o*rso=(s1 0:s2+d4=(1-)ood*>0s)](s)0*+oo=(0*:d./1)]." // find all factors
//var program = "[d2%(3*1+:2/)d1>]"         // collatz conjecture
//var program = "oo/+2/" // approximate sqrt

func init() {
    getopt.Flag(&debug, 'd', "Show debug output")
    getopt.FlagLong(&program, "--execute",'e',"Initialize a program")
}

func _stackstr(stack []int) string {
    var outstr string
    for _,val := range stack {
        outstr = fmt.Sprintf("%s %v",outstr,val)
    }
    return outstr
}

func _tos(stack []int) int {
    if len(stack) > 0 {
        return stack[len(stack)-1]
    } else {
        panic("One cannot peek at an empty stack")
    }
}

func _push(stack []int, val int) []int {
    return append(stack,val)
}

func _pop(stack []int) ([]int, int) {
    if len(stack) > 0 {
        var tos int
        tos, stack = stack[len(stack)-1], stack[:len(stack)-1]
        return stack, tos 
    } else {
        panic("Can't pop an empty stack")
    }
}

func _debug(mode int, stack []int, program string,ix int, returnstack []int)  {
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
        fmt.Printf("%s (%s ) %s{%s}%s %d (r: %s)\n", string(modestr[0]), _stackstr(stack), program[:ix],string(program[ix]),program[ix+1:],ix,_stackstr(returnstack))
    }
}


// The actual bytecodes
// each takes and returns a stack
func dup(stack []int) []int {
    defer func () {
        if r := recover(); r != nil {
            panic(fmt.Sprintf("%v\nCan't dup"))
        }
    }()
    stack = _push(stack,_tos(stack))
    return stack
}

func ovr(stack []int) []int {
    var a,b int
    stack, a = _pop(stack)
    stack, b = _pop(stack)
    stack = _push(stack,b)
    stack = _push(stack,a)
    return _push(stack,b)
}

func swp(stack []int) []int {
    var a,b int
    stack, a = _pop(stack)
    stack, b = _pop(stack)
    stack = _push(stack,a)
    return _push(stack,b)
}

func rol(stack []int) []int {
    var a,b,c int
    stack, a = _pop(stack)
    stack, b = _pop(stack)
    stack, c = _pop(stack)
    stack = _push(stack,b)
    stack = _push(stack,a)
    return _push(stack,c)
}

func dot(stack []int) []int {
    stack, val := _pop(stack)
    fmt.Printf("%d\n", val)
    return stack
}

func add(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    return _push(stack,a+b)
}

func sub(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    return _push(stack,a-b)
}

func mul(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    return _push(stack,a*b)
}

func div(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    return _push(stack,a/b)
}

func mod(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    return _push(stack,a%b)
}

func lst(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    if a < b {
        return _push(stack,1)
    } else {
        return _push(stack,0)
    }
}

func gtt(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    if a > b {
        return _push(stack,1)
    } else {
        return _push(stack,0)
    }
}

func not(stack []int) []int {
    var a int
    stack, a = _pop(stack)
    if a == 0 {
        return _push(stack,1)
    } else {
        return _push(stack,0)
    }
}

func equ(stack []int) []int {
    var a,b int
    stack, b = _pop(stack)
    stack, a = _pop(stack)
    if a == b {
        return _push(stack,1)
    } else {
        return _push(stack,0)
    }
}

const BYTE = 0
const COND = 1
const NUMB = 2
const SKIP = 3

func main() {
    getopt.Parse()
    stack := []int{}
    mode := BYTE
    returnstack := []int{}
    modestack := []int{}
    for _,str := range getopt.Args() {
        val,err := strconv.Atoi(str)
        if err == nil {
            stack = append(stack,val)
        }else{
            panic(err)
        }
    }
    for ix := 0; ix < len(program) ;ix += 1 {
        bytecode := byte(program[ix])
        _debug(mode,stack,program,ix,returnstack)
        switch mode {
        case NUMB:
            switch {
            case bytecode >= '0' && bytecode <= '9':
                stack,acc := _pop(stack)
                acc = acc * 10 + int(bytecode - '0')
                stack = _push(stack,acc)
            default:
                mode = BYTE
                ix -= 1
            }
        case SKIP: //skip the wrong part of a conditional
            switch bytecode {
            case '(': // nesting - oh no
                      // push mode on the modestack
                modestack = _push(modestack,mode)
            case ':': // else
                mode = COND
            case ')': // end if
                      // pop mode from modestack
                modestack,mode = _pop(modestack)
            }
        default:
            switch bytecode {
            case 'd':
                stack = dup(stack)
            case 'o':
                stack = ovr(stack)
            case 'r':
                stack = rol(stack)
            case 's':
                stack = swp(stack)
            case '.':
                stack = dot(stack) 
            case '+':
                stack = add(stack) 
            case '-':
                stack = sub(stack) 
            case '*':
                stack = mul(stack) 
            case '/':
                stack = div(stack) 
            case '%':
                stack = mod(stack) 
            case '<':
                stack = lst(stack) 
            case '=':
                stack = equ(stack) 
            case '~':
                stack = not(stack) 
            case '>':
                stack = gtt(stack) 
            case '(': // if
                var condition int
                stack,condition = _pop(stack)
                modestack = _push(modestack,mode)
                if condition == 0 {
                    mode = SKIP
                }else{
                    mode = COND
                }
            case ':': // else
                mode = SKIP
            case ')': // then (end if)
                modestack,mode = _pop(modestack)
            case '[': // loop
                returnstack = _push(returnstack,ix)
            case ']': // loop
                var condition int
                stack,condition = _pop(stack)
                if condition == 0 {
                    returnstack,_ = _pop(returnstack)
                } else {
                    ix = _tos(returnstack)
                }
            case ' ': //noop
                
            default:
                switch {
                case bytecode >= '0' && bytecode <= '9':
                    stack = _push(stack,int(bytecode - '0'))
                    mode = NUMB
                default:
                    panic("Unknown bytecode: [" + string(bytecode) + "]") 
                }
            }
        }
    }
    fmt.Printf("Stack: %s\n",_stackstr(stack))
}
