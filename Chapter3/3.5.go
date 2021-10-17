// Problem: 最も小さい項目がトップに来るスタックを並び変えるスタックを書く。別のスタックを1つ使うことも可。
// スタック以外のデータ構造にスタック上のデータをコピーしてはいけません。スタックは以下の操作のみ使用可：
// push, pop, peek, isEmpty

// Idea: Hold the max value in a variable, push all others to another stack s2. Push back all the elements in s2 to s1, and push temp element to s2. Repeat.

package main

import "fmt"

type Stack struct {
	stack []int
}

func (s *Stack) len() int {
	return len(s.stack)
}

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

func (s *Stack) peek() int {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) isEmpty() bool {
	if (len(s.stack) == 0) {
		return true
	} else {
		return false
	}
}

func (s *Stack) sortStack() *Stack {
	s2 := new(Stack)
	length := len(s.stack)
	val := 0
	count := 1
	temp := s.pop()

	for (len(s2.stack) != length) {
		// fmt.Println(s, s2, temp, length, len(s.stack))
		val = s.pop()
		if (val > temp) {
			s2.push(temp)
			temp = val
		} else {
			s2.push(val)
		}
		if s.isEmpty() {
			for i := 0; count <= len(s2.stack); i++ {
				// fmt.Println(count, len(s2.stack))
				s.push(s2.pop())
			}
			count++
			s2.push(temp)
			temp = s.pop()
			if s.isEmpty() {
				s2.push(temp)
				return s2
			}
		}
	}
	return s2
}

func main() {
	s := new(Stack)
	s.push(3)
	s.push(4)
	s.push(2)
	s.push(-1)

	fmt.Println(s)
	// fmt.Println(s.peek())
	// fmt.Println(s.isEmpty())
	fmt.Println(s.sortStack())

}