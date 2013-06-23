package util

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}, size int) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		size = s.size
		return
	}
	size = 0
	return
}

func (s *Stack) Iter() <-chan *Element {
	ch := make(chan *Element)
	go func() {
		var current *Element
		for i := 1; i <= s.size; i++ {
			if i == 1 {
				current = s.top
			} else {
				current = current.next
			}
			ch <- current
		}
		close(ch)
	}()
	return ch
}
