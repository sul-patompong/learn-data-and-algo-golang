package doublylinkedlists

import (
	"fmt"
	"log"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (list *List) Init() {
	list.Head = nil
	list.Tail = nil
	list.Length = 0
}

func (list *List) Push(value int) {
	node := &Node{
		Value: value,
	}

	if list.Head == nil {
		list.Head = node
		list.Tail = node
		node.Prev = nil
		node.Next = nil
	} else {
		node.Prev = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}

	list.Length++
}

func (list *List) Pop() *Node {
	temp := list.Tail

	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil
		list.Length--
		return temp
	}

	// Point tail to one upper
	// Set upper.next to null
	// Set temp.prev to null
	// Return temp
	upper := list.Tail.Prev
	list.Tail = upper
	upper.Next = nil
	temp.Prev = nil
	list.Length--

	return temp

}

func (list *List) Unshift(value int) {
	if list.Head == nil {
		list.Push(value)
		list.Length++
		return
	}

	node := &Node{
		Value: value,
		Next:  list.Head,
	}

	list.Head.Prev = node
	list.Head = node
	list.Length++
}

func (list *List) Shift() {
	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil
		list.Length--
		return
	}

	newHead := list.Head.Next
	list.Head.Next = nil
	list.Head = newHead
	list.Length--
}

func (list *List) Get(index int) *Node {
	if list.Length < index {
		log.Panic("out of range")
	}

	temp := list.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}

	return temp
}

func (list *List) Set(index int, value int) {
	if list.Length < index {
		log.Panic("out of range")
	}

	list.Get(index).Value = value
}

func (list *List) Insert(index int, value int) {
	if list.Length < index {
		log.Panic("out of range")
	}

	if index == 0 && list.Length != 0 {
		list.Unshift(value)
		return
	}

	if index == list.Length {
		list.Push(value)
		return
	}

	currentNode := list.Get(index)
	newNode := &Node{
		Value: value,
		Next:  currentNode,
		Prev:  currentNode.Prev,
	}

	currentNode.Prev.Next = newNode
	currentNode.Prev = newNode
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

	remove := list.Get(index)

	remove.Prev.Next = remove.Next
	remove.Next.Prev = remove.Prev

	remove.Next = nil
	remove.Prev = nil
	list.Length--
}

func (list *List) Print() {

	if list.Length == 0 {
		fmt.Println("list is null")
	} else {
		node := list.Head

		for node != nil {
			fmt.Print(node.Value)
			fmt.Print(" -> ")
			node = node.Next
		}

		fmt.Print("null")
		fmt.Print("\t| Head: ", list.Head.Value, " Tail: ", list.Tail.Value, " Length: ", list.Length, "\n")
	}
}

func Play() {
	var list List

	list.Init()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Print()

	list.Pop()
	list.Print()

	list.Unshift(9)
	list.Print()

	list.Shift()
	list.Print()

	list.Push(18)
	list.Print()

	fmt.Println(list.Get(1).Value)

	list.Push(84)
	list.Print()
	list.Set(1, 66)
	list.Print()

	list.Insert(3, 5)
	list.Print()

	list.Remove(0)
	list.Print()
}
