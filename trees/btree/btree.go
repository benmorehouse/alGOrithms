package main

import(
	"errors"
	"math"
)

type BTree struct{
	Root *Node
	M int // M stands for the max height we will allow in this tree 
	// for now we will just say M is equal to five.
	L int // L stands for the amount of 
	// this then means for access we get o(log base M of N)
	//we choose M and L on the basis of the size of the items that are being stored
}

type Node struct{
	Priority int
	Right map[int]*Node
	Left map[int]*Node
	Type interface{} // this is the interface that we assign it to. Not used unless in production
}

func (this* BTree) Init(){ // initialization for average cases
	// returns a btree. For now, we are to assume that L and M are sufficient for our data set.
	// though maybe at the end we can write a program which determines how many expected data points are there to be
	this.Root = nil // this should always be nil
	this.M = 5
	this.L = 5
}

func (this* BTree) Insert(i interface{})(error){
	_ , err := i.(int) // type assertion for input into function..
	if err != true{
		return errors.New("Type indifference for insertion")
	}
	// input is of correct type but how should we implement the node itself
	return nil
}

func (this *BTree) Search(i interface{})(*Node,error){
	input, err := i.(int) // type assertion for input into function..
	if err != true{
		return nil,errors.New("Type mismatch for BTree.Search")
	}
	// look at the currentNode
	currentNode := this.Root
	// now we need to look and see if it exists in currentNode
	bufferNode := currentNode.Left[input]
	if bufferNode != nil{
		return currentNode.Left[input], nil // if we find it right away
	}

	for i:=0;i<this.M;i++{
		var min int
		minNode := currentNode
		for _ , node := range currentNode.Left{
			// we look through each node in the left map to find which left/right we should go to 
			// we can just keep absolute value between
			if node.Priority == input{
				return node, nil
			}else if math.abs(node.Priority - input) < min{
				minNode = node
				min = node.Priority
				continue
			}
			continue
		}

		if input < minNode.Priority{
			// then currentNode will
			if minNode.Left == nil{
				// then we add into the current
				return nil, nil
			}
			currentNode = minNode.Left
			continue
		}
		currentNode = minNode.Right
	}
	return nil, nil

}

func (this *BTree) SearchHelper(input int)(*Node,error){

}
