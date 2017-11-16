#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "' pushes a literal from the program" {
  run ./rpnc -e "'9."
  assert_output "57"
}

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

