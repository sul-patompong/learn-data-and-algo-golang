package LinkedList

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (list *List) New(value int) *List {
	var node = &Node{
		Value: value,
	}

	list.Head = node
	list.Tail = node
	list.Length = 1

	return list
}

func (list *List) Push(value int) {
	var node = &Node{
		Value: value,
	}

	if list.Head == nil {
		list.Head = node
		list.Tail = node
		list.Length = 1
	} else {
		list.Tail.Next = node
		list.Tail = node
	}

	list.Length++
}

func (list *List) Pop() {

	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil

	}

	if list.Head != nil {
		var temp, previous *Node

		temp = list.Head

		for temp.Next != nil {
			previous = temp
			temp = temp.Next
		}

		list.Tail = previous
		previous.Next = nil

	}

	list.Length--
}

func (list *List) UnShift(value int) {
	if list.Head == nil {
		list.Push(value)
	} else {
		node := &Node{Value: value, Next: list.Head}
		list.Head = node

	}

	list.Length++
}

func (list *List) Shift() int {
	var shifted int
	if list.Head == nil {
		log.Panic("list is nil, you can't shift it")

	} else {
		shifted = list.Head.Value

		if list.Head.Next != nil {
			list.Head = list.Head.Next
		} else {
			list.Head = nil
			list.Tail = nil
		}

		list.Length--

	}

	return shifted
}

func (list *List) Get(index int) (int, error) {
	if list.Head == nil {
		log.Panic("list is nil")
	} else if index > list.Length {
		log.Panic("given index is larger than list length")
	} else if list.Length == index {
		return list.Tail.Value, nil
	} else {
		temp := list.Head

		for i := 0; i < index; i++ {
			temp = temp.Next
		}

		return temp.Value, nil
	}

	return 0, errors.New("not found")
}

func (list *List) Set(index int, value int) {
	if list.Head == nil {
		log.Panic("list is nil")
	} else if index > list.Length {
		log.Panic("given index is larger than list length")
	} else if list.Length == index {
		list.Tail.Value = value
	} else {
		temp := list.Head

		for i := 0; i < index; i++ {
			temp = temp.Next
		}

		temp.Value = value
	}
}

func (list *List) Insert(index int, value int) {
	if index == 0 {
		list.UnShift(value)
		return
	}

	if index == list.Length {
		list.Push(value)
		return
	}

	temp := list.Head
	var previous *Node

	for i := 0; i < index; i++ {
		previous = temp
		temp = temp.Next
	}

	node := &Node{
		Value: value,
		Next:  temp,
	}

	previous.Next = node
	list.Length++

}

func (list *List) Remove(index int) {
	if index == 0 {
		list.Shift()
		return
	}

	if index == list.Length-1 {
		list.Pop()
		return
	}

	temp := list.Head
	var previous *Node
	for i := 0; i < index; i++ {
		previous = temp
		temp = temp.Next
	}

	previous.Next = temp.Next
	list.Length--
}

func (list *List) Reverse() {
	// Head be Tail
	// Tail be Head
	//            T                   H
	//*      p    t    n
	//            1 -> 2 -> 3 -> 4 -> 5 -> null
	//*      p <- t    n
	//    null <- 1    2 -> 3 -> 4 -> 5 -> null
	//*           p <- t    n
	//            1 <- 2 -> 3 -> 4 -> 5 -> null

	// Head = 5 how to find 4
	// Set Head to tail
	// Itteration until 5 then use before value
	// set head.next to before

	head := list.Head
	list.Head = list.Tail
	list.Tail = head

	temp := list.Tail

	var previous, next *Node

	for i := 0; i < list.Length; i++ {
		next = temp.Next
		temp.Next = previous
		previous = temp
		temp = next
	}

}

func (list *List) Print() {

	node := list.Head

	for node != nil {
		fmt.Print(node.Value)
		fmt.Print(" -> ")
		node = node.Next
	}

	fmt.Print("null")
	fmt.Print("\t| Head: ", list.Head.Value, " Tail: ", list.Tail.Value, " Length: ", list.Length, "\n")
}

func Play() {
	var list List
	list.New(1)
	list.Push(2)
	list.Print()
	list.Push(3)
	list.Print()
	list.Push(4)
	list.Print()
	list.Push(5)
	list.Print()

	list.Pop()
	list.Print()
	list.Pop()
	list.Print()

	list.UnShift(-1)
	list.Print()

	list.Shift()
	list.Print()

	fmt.Println(list.Get(0))

	list.Set(2, 27657889)
	list.Print()

	list.Insert(0, 88)
	list.Print()

	list.Insert(3, 999)
	list.Print()

	list.Remove(2)
	list.Print()

	list.Reverse()
	list.Print()
}
