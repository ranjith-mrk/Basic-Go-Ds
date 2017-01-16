package main

import "fmt"
import "../stack"
import "strconv"

func isOperator(val string) bool{
  switch val {
    case "+", "-", "*", "/":
      return true
    default:
      return false
  }
}

func performOperation(val1, val2 int, operator string) int {
  switch operator {
    case "+":
      return val1 + val2
    case "-":
      return val1 - val2
    case "*":
      return val1 * val2
    case "/":
      return val1/val2
  }
  return 0
}

func evaluate(equation string) {
  s := stack.New()
  for _, k := range equation {
    c := string(k)
    if isOperator(c){
      b, a := s.Pop().(int), s.Pop().(int)
      val := performOperation(a, b, c)
      s.Push(val)
    }else{
      i, _ := strconv.Atoi(c)
      s.Push(i)
    }
  }
  fmt.Println(s.Pop().(int))
}

func main() {
  // equation := "231*+9-"
  equation := "234+*6-"
  evaluate(equation)
}