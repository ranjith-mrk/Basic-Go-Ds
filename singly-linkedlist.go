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

func (l *List) insertValue(val interface{}, at *Element) *Element {
  return l.insert(&Element{Value: val}, at)
}

func (l *List) insertAtEnd(val interface{}) *Element{
  var last = &l.root
  for i := 0; i < l.len - 1; i = i + 1 {
    last = last.next
  }
  return l.insertValue(val, last)
}

func (l *List) insertAtFront(val interface{}) *Element{
  return l.insertValue(val, &l.root)
}

func New() *List { return new(List).init() }

func main() {
  l := New()
  e := l.insertValue(1, l.root.next)
  l.insertAtFront(0)
  // l.insertAtEnd(2)
  // e1 := l.insertValue(3, e)
  // e2 := l.insertValue(2, e1)
  fmt.Println(e)
  fmt.Println(l.root.next.next.Value)
}