package main

import "fmt"

//First basic attempt at Stack and Queue to learn GO

type Node struct {
	name string
	next *Node
}

type Stack struct {
	name   string
	top    *Node
	bottom *Node
	next   *Node
}

type Queue struct {
	name string
	head *Node
	tail *Node
	next *Node
}

func stackPop(list *Stack) Node {
	if list.top != nil {
		node := *list.top
		list.top = list.top.next
		return node
	} else {
		return Node{name: "Pseudo Node"}
	}
}

func stackPush(node *Node, list *Stack) {
	if list.top == nil {
		list.top = node
		list.bottom = node
	}
	node.next = list.top
	list.top = node
}

func queuePush(node *Node, queue *Queue) {
	if queue.head == nil {
		queue.head = node
		queue.tail = node
	}

	queue.tail.next = node
	queue.tail = node
}

func queuePop(queue *Queue) Node {
	if queue.head != nil {
		node := *queue.head
		queue.head = queue.head.next
		return node
	} else {
		return Node{name: "Pseudo Node"}
	}

}

// No testing framework used yet
func testStack() {
	node1 := &Node{name: "Node 1"}
	node2 := &Node{name: "Node 2"}
	node3 := &Node{name: "Node 3"}

	stack1 := &Stack{name: "First list"}
	stackPush(node1, stack1)
	stackPush(node2, stack1)
	stackPush(node3, stack1)
	poppedVal1 := stackPop(stack1)
	fmt.Printf("The value of the popped should be Node3: %s\n", poppedVal1.name)
	poppedVal2 := stackPop(stack1)
	fmt.Printf("The value of the popped should be Node2: %s\n", poppedVal2.name)
	poppedVal3 := stackPop(stack1)
	fmt.Printf("The value of the popped should be Node1: %s\n", poppedVal3.name)

}

// No testing framework used yet
func testQueue() {
	node1 := &Node{name: "Node 1"}
	node2 := &Node{name: "Node 2"}
	node3 := &Node{name: "Node 3"}

	queue1 := &Queue{name: "Queue 1"}
	queuePush(node1, queue1)
	queuePush(node2, queue1)
	queuePush(node3, queue1)
	queuePush(&Node{name: "$"}, queue1)
	queuePush(&Node{name: "5"}, queue1)
	poppedVal1 := queuePop(queue1)
	fmt.Printf("The value of the popped should be Node1: %s\n", poppedVal1.name)
	poppedVal2 := queuePop(queue1)
	fmt.Printf("The value of the popped should be Node2: %s\n", poppedVal2.name)
	poppedVal3 := queuePop(queue1)
	fmt.Printf("The value of the popped should be Node3: %s\n", poppedVal3.name)
	poppedVal4 := queuePop(queue1)
	fmt.Printf("The value of the popped should be Node$: %s\n", poppedVal4.name)

}

func main() {
	fmt.Println("Testing Stack implementation")
	testStack()
	fmt.Println("")
	fmt.Println("Testing Queue implementation")
	testQueue()
}
