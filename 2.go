package main

import "fmt"

type Node struct {
	node *Node
	val  int
}

func (n *Node) Append(data int) {
	newNode := Node{nil, data}
	iter := n
	for iter.node != nil {
		iter = iter.node
	}
	iter.node = &newNode
}

func (n *Node) PrintNodes() {
	for n.node != nil {
		fmt.Println(n.val)
		n = n.node
	}
	fmt.Println(n.val)
}

func (n *Node) PartNodes(x int) {

	for n.node != nil {
		if n.val < x {
			temp := n.node.val
			n.node.val = n.val
			n.val = temp
		}
		n = n.node
	}
	n.PrintNodes()
	// return n
}

func main() {
	n := Node{}
	n.Append(3)
	n.Append(5)
	n.Append(8)
	n.Append(5)
	n.Append(10)
	n.Append(2)
	n.Append(1)

	n.PartNodes(5)
	fmt.Println(n.val)
	// n.PrintNodes()

}
