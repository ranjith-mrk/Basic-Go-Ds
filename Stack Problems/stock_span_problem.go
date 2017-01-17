//Stock span problem

//http://www.geeksforgeeks.org/the-stock-span-problem/

package main

import "../stack"
import "fmt"

func calculateSpan(prices ...int) {
  s := stack.New()
  s.Push(0)
  count := len(prices)
  list := make([]int, count)
  list[0] = 1
  for i :=1; i < count; i++{
    for(!s.IsEmpty() &&  prices[s.Peek().(int)] <= prices[i]){
      s.Pop()
    }
    if(s.IsEmpty()){
      list[i] = i+1
    }else{
      list[i] = i - s.Peek().(int)
    }
    s.Push(i)
  }
  fmt.Println(list)
}

func main() {
  // prices := []int{11, 13, 21, 3, 44}
  prices := []int{10, 4, 5, 90, 120, 80}
  calculateSpan(prices...)
}