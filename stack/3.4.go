package stack

import "fmt"

type Stack struct {
	stack []int
}

type MyQueue struct {
	stack1 Stack
	stack2 Stack
}

func (s *Stack) len() int {
	return len(s.stack)
}

// func (s *Stack) append(i) *Stack {
// 	s = append(s.stack, i)
// 	return s
// }

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (s *Stack) push(i int) *Stack {
	// var s.stack []int
	s.stack = append(s.stack, i)
	return s
}

func (s *Stack) pop() int {
	popped := s.stack[len(s.stack)-1]
	s.stack = remove(s.stack, len(s.stack)-1)
	return popped
}

func (q *MyQueue) q_push(i int) *MyQueue {
	q.stack1.push(i)
	return q
}

func (q *MyQueue) q_pop() int {
	// fmt.Println(q.stack1.len())
	if q.stack1.len() == 1 {
		q_popped := q.stack1.pop()
		return q_popped
	}
	length := q.stack1.len()
	for i := 0; i < length; i++ {
		// fmt.Println(i)
		q.stack2.push(q.stack1.pop())
	}
	// fmt.Println(q.stack2)
	q_popped := q.stack2.pop()
	length2 := q.stack2.len()
	for i := 0; i < length2; i++ {
		q.stack1.push(q.stack2.pop())
	}
	return q_popped
}

func main() {
	s := new(MyQueue)
	s.q_push(1)
	s.q_push(4)
	s.q_push(-3)
	s.q_push(10)
	fmt.Println(s)
	fmt.Println(s.q_pop(), s)
	fmt.Println(s.q_pop(), s)
	fmt.Println(s.q_pop(), s)
	fmt.Println(s.q_pop(), s)
}
