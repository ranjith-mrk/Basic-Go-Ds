package main

import "fmt"

type Element struct {
  next *Element
  list *List
  Value interface{}
}

type List struct {
  root Element
  len int
}

func (e *Element) Next() *Element{
  if n := e.next; e.list != nil && n != &e.list.root{
    return n
  }
  return nil
}

func (l *List) init() *List{
  l.root.next = &l.root
  l.len = 0
  return l
}

func (l *List) insert(e, at *Element) *List {
  n := at.Next()
  at.next = e
  e.next = n
  e.list = l 
  l.len = l.len + 1
  return l
}


func main() {
  l := new(List).init()
  fmt.Println(l)
}