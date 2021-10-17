package main

import "fmt"

type ThreeStack struct {
	stack []int
	mark  int
}

func (s *ThreeStack) push(num int, x int) *ThreeStack {
	counter := 0
	for i := 0; i < len(s.stack); i++ {
		if s.stack[i] == s.mark {
			counter++
			if counter == num {
				s.stack.splice(i, 0, x)
				break
			}
		}
	}
	return s
}

func (s *ThreeStack) pop(num int, i int) int {
	counter := 0
	popped := -1
	for i := 0; i < len(s.stack); i++ {
		if s.stack[i] == s.mark {
			counter++
			if counter == num {
				popped := s.stack[len(s.stack)-1]
				s.stack = remove(s.stack, len(s.stack))
				return popped
			}
		}
	}
	return popped
}

func main() {
	s := new(ThreeStack)
	s.push(1, 1)
	s.push(1, 4)
	s.push(2, -3)
	s.push(2, 10)
	fmt.Println(s)
}
