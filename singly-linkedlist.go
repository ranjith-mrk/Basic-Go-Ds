package list

type Element struct {
  next *Element
  list *List
  Value interface{}
}

type List struct {
  root Element
  len int
}

func (e *Element) next() {
  if n := e.next(); e.list != nil && n != &e.list.root{
    return n
  }
  return nil
}


func main() {
  
}