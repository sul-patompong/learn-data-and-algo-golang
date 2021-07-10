package binarytree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) Init() {
	b.root = nil
}

func (b *BinaryTree) Insert(value int) {
	newNode := &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}

	if b.root == nil {
		b.root = newNode
		return
	}

	temp := b.root

	for temp != nil {
		fmt.Println("Node: ", temp)
		if temp == nil {
			temp = newNode
			return
		}

		if value < temp.Value {
			if temp.Left == nil {
				temp.Left = newNode
				return
			} else {
				temp = temp.Left
			}
		} else if value > temp.Value {
			if temp.Right == nil {
				temp.Right = newNode
				return
			} else {
				temp = temp.Right
			}
		}
	}

	// check the value is less or greather
	// if less go to left
	// if greather fo ro right
	// if leaf node insert it

}

func (b *BinaryTree) Contain(value int) bool {
	if b == nil {
		return false
	}

	temp := b.root

	for temp != nil {
		fmt.Println("Search at: ", temp)
		if temp.Value == value {
			return true
		}

		if value < temp.Value {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}

	return false
}

// func (n *Node) Find(value int) bool {

// }

func Play() {
	var b BinaryTree

	b.Init()
	b.Insert(47)
	b.Insert(76)
	b.Insert(21)
	b.Insert(18)
	b.Insert(52)

	fmt.Println(b.root)
	fmt.Println("\t", b.root.Left)
	fmt.Println("\t\t", b.root.Left.Left)

	fmt.Println("\t", b.root.Right)
	fmt.Println("\t\t", b.root.Right.Left)

	fmt.Println(b.Contain(47))
}
