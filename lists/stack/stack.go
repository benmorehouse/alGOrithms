package stack

type Node struct{
	Data int
	Next *Node
}

type Stack struct{
	Array []*Node
}

func (this* Stack)Push(i interface{}){
	input, err := i.(int)
	if err != true{
		return
	}
	var newNode *Node
	newNode.Data = input
	if len(this.Array) == 0{
		// we are trying to make a new node in this function
		// and in order to do this we need to:
		// make the node with the input
		// and we need to update this (which is of type stack) to now have a new head!
		newNode.Next = nil
		this.Array = append(this.Array,newNode)
		return
	}else{
		newNode.Next = nil
		this.Array[len(this.Array)-1].Next = newNode
		this.Array = append(this.Array,newNode)
		return
	}
}

func (this *Stack) Peek() (*Node){
	if len(this.Array) == 0{
		return nil
	}
	return this.Array[len(this.Array)-1]
}

func (this *Stack) Pop(){
	if len(this.Array) == 0{
		return
	}
	this.Array = this.Array[:len(this.Array)-1]
	return
}
