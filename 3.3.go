package main

import "fmt"

type SetofStacks struct {
	set_of_stacks [][]int
	stack         []int
	limit         int
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func set_remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

func (s *SetofStacks) push(i int) *SetofStacks {
	s.stack = append(s.stack, i)
	if len(s.stack) == s.limit {
		s.set_of_stacks = append(s.set_of_stacks, s.stack)
		s.stack = nil
	}
	return s
}

func (s *SetofStacks) pop() int {
	if (len(s.stack) == 0) && (len(s.set_of_stacks) > 0) {
		popped := s.set_of_stacks[len(s.set_of_stacks)-1][s.limit-1]
		s.stack = remove(s.set_of_stacks[len(s.set_of_stacks)-1], s.limit-1)
		s.set_of_stacks = set_remove(s.set_of_stacks, len(s.set_of_stacks)-1)
		return popped
	} else {
		popped := s.stack[len(s.stack)-1]
		s.stack = remove(s.stack, len(s.stack)-1)
		return popped
	}
}

func main() {
	s := new(SetofStacks)
	s.limit = 3
	s.push(1)
	s.push(4)
	s.push(-3)
	s.push(10)
	s.push(20)
	fmt.Println(s.stack, s.set_of_stacks, len(s.set_of_stacks), len(s.stack))
	fmt.Println(s.pop())
	fmt.Println(s.stack, s.set_of_stacks, len(s.set_of_stacks), len(s.stack))
}
