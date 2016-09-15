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

func (l *List) insert(e, at *Element) *Element {
  n := at.Next()
  at.next = e
  e.next = n
  e.list = l 
  l.len = l.len + 1
  return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
  return l.insert(&Element{Value: v}, at)
}

func New() *List { return new(List).init() }

func main() {
  l := New()
  e := l.insertValue(1, l.root.next)
  e1 := l.insertValue(3, e)
  e2 := l.insertValue(2, e1)
  fmt.Println(e2)
  fmt.Println(l.root.next.next.next.Value)
}