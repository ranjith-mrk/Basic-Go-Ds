//Minimum number of bracket reversals needed
//http://www.geeksforgeeks.org/minimum-number-of-bracket-reversals-needed-to-make-an-expression-balanced/

package main 

import "fmt"
import "../stack"
import "strings"
import "math"

func evaluateExpression(expression string) {
  s := stack.New()
  exp := strings.Split(expression, "")
  for _, i := range exp {
    if i == "{" {
      s.Push(i)
    }else if (!s.IsEmpty() && s.Peek().(string) == "{")  {
      s.Pop()
    }else{
      s.Push(i)
    }
  }
  var open, closed float64
  for !s.IsEmpty() {
    if s.Peek().(string) == "{" {
      open++
    }else{
      closed++
    }
    s.Pop()
  }
  total := math.Ceil(open/2) + math.Ceil(closed/2)
  fmt.Println(total)
}

func main() {
  // expression := "}}{{"
  expression := "}{{}}{{{"
  evaluateExpression(expression)
}