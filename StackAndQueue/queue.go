package stackandqueue

import "fmt"

type Queue struct {
	First  *Node
	Last   *Node
	Length int
}

func (q *Queue) Init(value int) {
	q.First = &Node{
		Value: value,
		Next:  nil,
	}

	q.Last = q.First
	q.Length = 1
}

func (q *Queue) Push(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	q.Last.Next = newNode
	q.Last = newNode
	q.Length++
}

func (q *Queue) Shift() *Node {
	first := q.First

	q.First = first.Next
	first.Next = nil

	q.Length--
	return first
}

func PlayQueue() {
	var q Queue
	q.Init(1)
	q.Print()

	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Print()

	q.Shift()
	q.Print()

	q.Shift()
	q.Print()
}

func (q *Queue) Print() {

	if q.Length == 0 {
		fmt.Println("list is null")
	} else {
		node := q.First

		for node != nil {
			fmt.Print(node.Value)
			fmt.Print(" -> ")
			node = node.Next
		}

		fmt.Print("null")
		fmt.Print("\t| Head: ", q.First.Value, " Length: ", q.Length, "\n")
	}
}
