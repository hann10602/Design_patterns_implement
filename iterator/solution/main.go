package main

import "fmt"

type Follower interface {
	Receive(message string)
}

type Profile struct {
	name string
}

func (p Profile) Receive(message string) {
	fmt.Printf("%s: %s\n", p.name, message)
}

var arrayOfFollowers = []Follower{
	Profile{name: "Peter"},
	Profile{name: "Mary"},
	Profile{name: "Alan"},
}

type LinkedNode struct {
	val Follower
	next *LinkedNode
}

var linkedListOfFollowers = &LinkedNode{
	val: Profile{name: "Perter"},
	next: &LinkedNode{
		val: Profile{name: "Mary"},
		next: &LinkedNode{
			val: Profile{name: "Alan"},
			next: nil,
		},
	},
}

type TreeNode struct {
	val Follower
	children []TreeNode
}

var treeOfFollowers = &TreeNode{
	val: Profile{name: "Peter"},
	children: []TreeNode{
		{
			val: Profile{name: "Mary"},
		},
		{
			val: Profile{name: "Alan"},
		},
		{
			val: Profile{name: "Alex"},
			children: []TreeNode{
				{
					val: Profile{name: "Vicky"},
				},
				{
					val: Profile{name: "Alice"},
				},
			},
		},
	},
}

func sendMessageForArray(msg string) {
	for i := range arrayOfFollowers {
		arrayOfFollowers[i].Receive(msg)
	}
}

func sendMessageForLinkedList(msg string) {
	node := linkedListOfFollowers

	for node != nil {
		node.val.Receive(msg)
		node = node.next
	}
}

func sendMessageForTree(node *TreeNode, msg string) {
	if node == nil {return}

	node.val.Receive(msg)

	for i := range node.children {
		sendMessageForTree(&node.children[i], msg)
	}
}

func main() {
	msg := "Hello world"

	fmt.Println("Sending message for array")
	sendMessageForArray(msg)
	fmt.Println("Sending message for linked list")
	sendMessageForLinkedList(msg)
	fmt.Println("Sending message for tree")
	sendMessageForTree(treeOfFollowers, msg)
}