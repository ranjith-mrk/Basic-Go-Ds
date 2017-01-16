//Next Greater Element
//http://www.geeksforgeeks.org/next-greater-element/
package main

import "../stack"
import "fmt"

func findNextGreatest(elements ...int) {
  s := stack.New()
  noMapping := make(map[int]int)
  s.Push(elements[0])
  for _, currentElemet := range elements[1:] {
    for (!s.IsEmpty() && (currentElemet > s.Peek().(int))){
      noMapping[s.Pop().(int)] = currentElemet
    }
    if(s.IsEmpty() || (currentElemet < s.Peek().(int))) {
      s.Push(currentElemet)
    }
  }
  for !s.IsEmpty(){
    noMapping[s.Pop().(int)] = -1
  }
  fmt.Println(noMapping)
}

func main() {
  elements := []int{11, 13, 21, 3, 44}
  findNextGreatest(elements...)
}