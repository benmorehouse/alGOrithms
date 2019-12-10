package queue

import(
)

type Node struct{
	Data int
	Next *Node
}

type Queue struct{
	Array []*Node
}

func (this* Queue)Push(i interface{}){
	input, err := i.(int)
	if err != true{
		return
	}
	newNode := Node{
		Data: input,
		Next: nil,
	}
	if len(this.Array) == 0{
		// we are trying to make a new node in this function
		// and in order to do this we need to:
		// make the node with the input
		// and we need to update this (which is of type stack) to now have a new head!
		newNode.Next = nil
		this.Array = append(this.Array,&newNode)
		return
	}else{
		newNode.Next = nil
		this.Array[len(this.Array)-1].Next = &newNode
		this.Array = append(this.Array,&newNode)
		return
	}
}

func (this *Queue) Peek() (*Node){
	if len(this.Array) == 0{
		return nil
	}
	return this.Array[0]
}

func (this *Queue) Pop(){
	if len(this.Array) == 0{
		return
	}
	this.Array = this.Array[1:]
	return
}
