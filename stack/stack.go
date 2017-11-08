package stack

import "fmt"

type Stack struct {
    s []int
}

func New() *Stack {
    return new(Stack)
}

func (this *Stack) String() string {
    var outstr string = "("
    for _,val := range this.s {
        outstr = fmt.Sprintf("%s %v",outstr,val)
    }
    return fmt.Sprintf("%s )",outstr)
}

func (this *Stack) Push(value int){
    this.s = append(this.s,value)
}

func (this *Stack) Pop() int {
    length := len(this.s)
    value := this.s[length-1]
    this.s = this.s[:length-1]
    return value
}

func (this *Stack) Peek() int {
    length := len(this.s)
    return this.s[length-1]
}
