package hashtable

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	h := &HashTable{}

	for i := range h.array {
		h.array[i] = &bucket{}
	}

	return h
}

func (h *HashTable) Insert(key string) {
	hashed := hash(key)

	fmt.Println("The hashed number => ", hashed)
	h.array[hashed].Insert(key)
}

func (h *HashTable) Search(key string) bool {
	return h.array[hash(key)].search(key)
}

func (h *HashTable) Delete(key string) {
	h.array[hash(key)].delete(key)
}

func (b *bucket) Insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{
			key: key,
		}

		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("the key is already existed")
	}
}

func (b *bucket) search(key string) bool {
	current := b.head

	for current != nil {
		if current.key == key {
			return true
		}

		current = current.next
	}

	return false
}

func (b *bucket) delete(key string) {
	var previous *bucketNode
	current := b.head
	next := b.head.next

	if b.head.key == key {
		current.next = nil
		b.head = next
		return
	}

	for current != nil {
		if current.key == key {
			previous.next = current.next
			current.next = nil
		}

		previous = current
		current = current.next
	}

}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func Play() {
	h := Init()

	fmt.Println(h)

	h.Insert("Beam")
	h.Insert("Mui")
	h.Insert("Boo")
	h.Insert("Eak")
	h.Insert("Dumm")
	h.Insert("Kob")
	h.Insert("Soey")
	h.Insert("Jaree")

	fmt.Println(h.Search("Dumm"))

	h.Delete("Kob")

}
