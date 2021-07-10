package main

import (
	"fmt"
	"os"
	"os/exec"

	binarytree "github.com/sul-patompong/learn-go-data-and-algo/BinaryTree"
	doublylinkedlists "github.com/sul-patompong/learn-go-data-and-algo/DoublyLinkedLists"
	graphs "github.com/sul-patompong/learn-go-data-and-algo/Graphs"
	hashtable "github.com/sul-patompong/learn-go-data-and-algo/HashTable"
	LinkedList "github.com/sul-patompong/learn-go-data-and-algo/LinkedLists"
	stackandqueue "github.com/sul-patompong/learn-go-data-and-algo/StackAndQueue"
)

func main() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
	fmt.Println("Listked list")
	LinkedList.Play()
	fmt.Println("\n\nDoubly Listked list")
	doublylinkedlists.Play()
	fmt.Println("\n\nStack")
	stackandqueue.Play()
	fmt.Println("\n\nQueue")
	stackandqueue.PlayQueue()
	fmt.Println("\n\nBinary Tree")
	binarytree.Play()
	fmt.Println("\n\nGraph")
	graphs.Play()
	fmt.Println("\n\nHash Table")
	hashtable.Play()
}
