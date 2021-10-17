package main

import "fmt"

type Stack struct {
	stack   []int
	counter int
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (s *Stack) push(i int) *Stack {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, i)
		s.counter = 1
	} else if i > s.stack[len(s.stack)-1] {
		s.stack = append([]int{i}, s.stack...)
		s.counter = 0
	} else {
		s.stack = append(s.stack, i)
		s.counter = 1
	}
	return s
}

func (s *Stack) pop() int {

	if s.counter == 0 {
		popped := s.stack[0]
		s.stack = remove(s.stack, 0)
		s.counter = 1
		return popped
	} else {
		popped := s.stack[len(s.stack)-1]
		s.stack = remove(s.stack, len(s.stack)-1)
		return popped
	}
}

func (s *Stack) min() int {
	return s.stack[len(s.stack)-1]
}

func main() {
	s := new(Stack)
	s.push(1)
	s.push(4)
	s.push(-3)
	s.push(10)
	fmt.Println(s, s.min(), s.pop())
	fmt.Println(s, s.min(), s.pop())
	fmt.Println(s, s.min(), s.pop())
}
