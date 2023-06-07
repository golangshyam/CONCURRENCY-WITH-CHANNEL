package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{value: value, next: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func SeparateEvenOdd(ll *LinkedList) (*LinkedList, *LinkedList) {
	evenList := &LinkedList{}
	oddList := &LinkedList{}

	current := ll.head

	for current != nil {
		if current.value%2 == 0 {
			evenList.Insert(current.value)
		} else {
			oddList.Insert(current.value)
		}
		current = current.next
	}

	return evenList, oddList
}

func PrintList(ll *LinkedList) {
	current := ll.head

	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}

	fmt.Println()
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	list := &LinkedList{}

	for _, num := range numbers {
		list.Insert(num)
	}

	fmt.Println("Original List:")
	PrintList(list)

	evenList, oddList := SeparateEvenOdd(list)

	fmt.Println("Even List:")
	PrintList(evenList)

	fmt.Println("Odd List:")
	PrintList(oddList)
}