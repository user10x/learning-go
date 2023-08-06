package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewNode(data int) *Node {
	return &Node{data: data, next: nil}
}

func (l *LinkedList) addToFrontLinkedList(node *Node) {
	if node == nil {
		return
	}
	if l.head == nil {
		l.head = node
		return
	}

	node.next = l.head
	l.head = node
}

func (l *LinkedList) traverLinkedList() {
	current := l.head
	for current != nil {
		fmt.Printf("data is %d\n", current.data)
		current = current.next
	}
}

func detectLoop(l *LinkedList) bool {
	slowNode := l.head
	fastNode := l.head.next
	for slowNode != nil && fastNode != nil {
		if slowNode.data == fastNode.data {
			fmt.Println("cycle at ", slowNode.data)
			return true
		}
		if fastNode.next == nil {
			return false
		}
		slowNode = slowNode.next
		fastNode = fastNode.next.next
	}
	return false
}

func main() {

	l := NewLinkedList()
	l.addToFrontLinkedList(NewNode(4))
	l.addToFrontLinkedList(NewNode(3))
	l.addToFrontLinkedList(NewNode(2))
	l.addToFrontLinkedList(NewNode(1))
	l.traverLinkedList()

	fmt.Println(detectLoop(l))

}
