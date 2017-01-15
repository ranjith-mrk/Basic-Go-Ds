//Infix to Postfix

package main

import "fmt"
import "../stack"
import "bytes"

func operatorPriority(operator string) int {
  switch operator {
    case "+", "-":
      return 1
    case "*", "/":
      return 2
    case "^":
      return 3
  }
  return -1
}

func canInsertInStack(operator1, operator2 string) bool{
  return operatorPriority(operator2) > operatorPriority(operator1)
}

func PrintPostfix(equation string) {
  s := stack.New()
  var buffer bytes.Buffer
  for _, k := range equation{
    c := string(k)
    if c == "(" {
        s.Push(c)
      }else if c == ")"{)
        for {
          if (s.IsEmpty() || s.Peek().(string) == "("){ 
            s.Pop()
            break
          }
          buffer.WriteString(s.Pop().(string))
        }
      }else if operatorPriority(c) == -1 {
        buffer.WriteString(c)
      } else if (s.IsEmpty() || (!s.IsEmpty() && canInsertInStack(s.Peek().(string), c))) {
        s.Push(c)
      } else if(!s.IsEmpty()){
        for {
          if (operatorPriority(s.Peek().(string)) >= operatorPriority(c)) {
            buffer.WriteString(s.Pop().(string))
            if (!s.IsEmpty()){
              continue
            }
          }
          break
        }
        s.Push(c)
      }
  }
  for !s.IsEmpty(){
    buffer.WriteString(s.Pop().(string))
  }
  fmt.Println(buffer.String())
}

func main(){
  equation := "a+b*c+d"  
  PrintPostfix(equation)
}

