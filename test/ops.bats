#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "d duplicated the top of the stack" {
  run go run rpnc.go -e 'd' 9
  assert_output 'Stack: ( 9 9 )'
}

@test "r rotates the top of the stack" {
  run go run rpnc.go -e 'r' 1 2 3
  assert_output 'Stack: ( 2 3 1 )'
}

@test "o pulls over the top of the stack" {
  run go run rpnc.go -e 'o' 1 2 3
  assert_output 'Stack: ( 1 2 3 2 )'
}

@test "s swaps the top of the stack" {
  run go run rpnc.go -e 's' 1 2 3
  assert_output 'Stack: ( 1 3 2 )'
}

@test ". pops the stack and prints the value" {
  result=$(go run rpnc.go -e '.' 1 2 3)
  echo $result | head -1 | grep '3'
  echo $result | tail -1 | grep 'Stack: ( 1 2 )'
}

			#case '+':
				#add(mainstack)
			#case '-':
				#sub(mainstack)
			#case '*':
				#mul(mainstack)
			#case '/':
				#div(mainstack)
			#case '%':
				#mod(mainstack)
				#// comparison
			#case '<':
				#lst(mainstack)
			#case '=':
				#equ(mainstack)
			#case '>':
				#gtt(mainstack)
				#// logical negation
			#case '~':
				#not(mainstack)
				#// conditional
			#case '(': // if
				#if mainstack.Pop() == 0 {
					#mode = SKIP
				#} else {
					#mode = COND
				#}
			#case ':': // else
				#mode = SKIP
			#case ')': // then (end if)
				#mode = modestack.Pop()
			#case '[': // loop
				#returnstack.Push(ix)
			#case ']': // loop
				#var condition int
				#condition = mainstack.Pop()
				#if condition == 0 {
					#_ = returnstack.Pop()
				#} else {
					#ix = returnstack.Peek()
				#}
				#// manipulate memory
			#case '^':
				#mainstack.Push(0)
			#case '@':
				#fet(mainstack)
			#case '!':
				#sto(mainstack)

			#case ' ': //noop

			#default:
				#switch {
				#case bytecode >= '0' && bytecode <= '9':
					#mainstack.Push(int(bytecode - '0'))
					#mode = NUMB
				#default:
					#panic("Unknown bytecode: [" + string(bytecode) + "]")
				#}

