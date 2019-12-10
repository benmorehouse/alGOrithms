package main

import(
	"errors"
)

type Node struct{ // a Node structure
	Priority int
}

type Heap struct{
	Array []*Node //an array to hold its priorities
}

func (this *Heap) Insert(i interface{}) error{
	input, ok := i.(int) // type assertion for input into function..
	if ok != true{
		return errors.New("Error: type assertion returned false")
	}

	newNode := Node{
		Priority: input,
	}

	this.Array = append(this.Array,newNode)
	this.TrickleUp(len(this.Array)-1)// we now trickle up the heap
}

func (this *Heap) Peek()(*Node){
	if len(this.Array) == 0{
		return
	}

	return this.Array[0]
}

func (this *Heap) TrickleUp(index int){
	if index < 0 || index >= len(this.Array){
		return
	}

}
