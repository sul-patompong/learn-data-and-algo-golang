package stackandqueue

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top    *Node
	Length int
}

func (s *Stack) Init(value int) {
	s.Top = &Node{
		Value: value,
		Next:  nil,
	}

	s.Length = 1
}

func (s *Stack) Push(value int) {
	newNode := &Node{
		Value: value,
		Next:  s.Top,
	}

	s.Top = newNode
	s.Length++
}

func (s *Stack) Pop() *Node {
	pop := s.Top
	s.Top = s.Top.Next
	s.Length--
	return pop
}

func Play() {
	var s Stack

	s.Init(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Print()

	fmt.Println(s.Pop())
	s.Print()
}

func (list *Stack) Print() {

	if list.Length == 0 {
		fmt.Println("list is null")
	} else {
		node := list.Top

		for node != nil {
			fmt.Print(node.Value)
			fmt.Print(" -> ")
			node = node.Next
		}

		fmt.Print("null")
		fmt.Print("\t| Head: ", list.Top.Value, " Length: ", list.Length, "\n")
	}
}
