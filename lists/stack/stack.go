package main

import(
	"fmt"
)

type node struct{
	data int
	next *node
}

type stack struct{
	head *node // and that is all we need for now
}

func (this* stack)Push(input int){
	var newNode *node
	newNode.data = input
	if this.head == nil{
		// we are trying to make a new node in this function
		// and in order to do this we need to:
		// make the node with the input
		// and we need to update this (which is of type stack) to now have a new head!
		newNode.next = nil
		this.head = newNode
		return
	}else{
		newNode.next = nil
		this.head.next = newNode
		this.head = newNode
		return
	}
}

func (this *stack) PeekData(){
	if this.head == nil{
		return
	}

	fmt.Println("Head node data is:",this.head.data)
}
