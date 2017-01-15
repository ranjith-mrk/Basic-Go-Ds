package stack

// import "fmt"

type Stack struct{
  stack []interface{}
  len   int
}

func New() *Stack{
  s := &Stack{}
  s.stack = make([]interface{}, 0)
  return s
}

func (s *Stack) Len() int {
  return s.len
}

func (s *Stack) Pop() (el interface{}) {
    el = s.stack[0]
    s.stack = s.stack[1:]
    s.len--
    return el
}

func (s *Stack) Push(el interface{}) {
  top := make([]interface{}, 1)
  top[0] = el
  s.stack = append(top, s.stack...)
  s.len++
}

func (s *Stack) Peek() (el interface{}) {
  return s.stack[0]
}

func (s *Stack) IsEmpty() bool {
  return s.len == 0
}