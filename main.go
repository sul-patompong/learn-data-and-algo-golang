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
	basicsorting "github.com/sul-patompong/learn-go-data-and-algo/Sorting/BasicSorting"
	mergesort "github.com/sul-patompong/learn-go-data-and-algo/Sorting/MergeSort"
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

	fmt.Println(basicsorting.BubbleSort([]int{3, 4, 5, 2, 1, 9, 5, 6, 7, 8, 7, 5, 4, -7, 3, 232, 3, 45, 56}))
	fmt.Println(basicsorting.SelectionSort([]int{3, 4, 5, 2, 1, 9, 5, 6, 7, 8, 7, 5, 4, -7, 3, 232, 3, 45, 56}))

	fmt.Println("Insertion sorted")
	fmt.Println(basicsorting.InsertionSort([]int{3, 4, 5, 2, 1, 9, 5, 6, 7, 8, 7, 5, 4, -7, 3, 232, 3, 45, 56}))

	fmt.Println("Merge sorted")
	fmt.Println(mergesort.MergeSort([]int{3, 4, 5, 2, 1, 9, 5, 6, 7, 8, 7, 5, 4, -7, 3, 2, 3, 5, 6}))

}
