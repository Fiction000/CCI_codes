package main

import (
	"container/list"
	"fmt"
)

func check_linked_palin(l *list.List) bool {
	front := l.Front()
	back := l.Back()
	i := 0

	for l.Len()/2 != i {
		fmt.Println(front.Value, back.Value)
		if front.Value != back.Value {
			return false
		}
		front = front.Next()
		back = back.Prev()
		i++
	}
	return true
}

func main() {
	l := list.New()
	l.PushBack("a")
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("a")
	l.PushBack("a")

	fmt.Println(check_linked_palin(l))

}
