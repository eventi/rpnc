package main

import "github.com/eventi/rpnc/stack"
import "fmt"

func stackstr(stack *stack.Stack) {
	fmt.Println(stack.Pop())
}

func main() {
	stack := stack.New()
	stack.Push(1)
	dup(stack)
	fmt.Println(stack.Pop())
	stackstr(stack)
}
