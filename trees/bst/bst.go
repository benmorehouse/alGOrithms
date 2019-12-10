package bst

import(
	"errors"
	"fmt"
)

type Node struct{
	Priority int
	Right *Node
	Left *Node
}

func (head *Node) Insert(i interface{}) error{
	input, err := i.(int) // type assertion for input into function..
	if err != true{
		return errors.New("Error: type assertion returned false")
	}

	newNode := Node{
		Priority: input,
	}

	if head == nil{
		head = &newNode
		return nil
	}

	currentNode := head
	for currentNode != nil{
		if input < currentNode.Priority{
			if currentNode.Left == nil{
				currentNode.Left = &newNode
				return nil
			}
			currentNode = currentNode.Left
		}else if input > currentNode.Priority{
			if currentNode.Right == nil{
				currentNode.Right = &newNode
				return nil
			}
			currentNode = currentNode.Right
		}else{
			return nil
		}
	}
	return nil
}

func (head *Node) Search(i interface{}) error{ // this should probably just return the node itself
	input, err := i.(int)
	if err != true{
		return errors.New("Error: type assertion returned false")
	}
	if head == nil{
		return errors.New("Root node is nil")
	}

	if input < head.Priority{
		if head.Left != nil{
			return head.Left.Search(input)
		}else{
			return errors.New("Node doesnt exists")
		}
	}else if input > head.Priority{
		if head.Right != nil{
			return head.Right.Search(input)
		}else{
			return errors.New("Node doesnt exists")
		}
	}else{
		return nil
	}
	return errors.New("Something wrong with search function")
}

func (head *Node) BFS(input int)error{
	var queue []*Node
	if head == nil{
		return errors.New("Entered an empty node")
	}else if head.Priority == input{
		return errors.New("Entered an empty node")
	}

	if head.Right == nil{
		queue = append(queue, head.Left)
	}else if head.Left == nil{
		queue = append(queue, head.Right)
	}else{
		queue = append(queue, head.Left)
		queue = append(queue, head.Right)
	}
	return BFSHelper(input, queue)
}

func BFSHelper(input int, queue []*Node) error{
	var newQueue []*Node
	if len(queue) == 0{
		return errors.New("couldnt find the node")
	}
	for i:=0;i<len(queue);i++{
		if queue[i].Priority == input{
			return nil
		}else{
			if queue[i].Right == nil && queue[i].Left == nil{
				continue
			}else if queue[i].Right == nil{
				newQueue = append(newQueue, queue[i].Left)
			}else if queue[i].Left == nil{
				newQueue = append(newQueue, queue[i].Right)
			}else{
				newQueue = append(newQueue, queue[i].Left)
				newQueue = append(newQueue, queue[i].Right)
			}
		}
	}
	return BFSHelper(input, newQueue)
}

func (head *Node) DFS(input int)error{
	if head == nil{
		return errors.New("Entered an empty node")
	}else if head.Priority == input{
		return nil
	}

	var stack []*Node
	stack = append(stack, head)
	c := make(chan bool, 2) // using a buffered channel so we dont get more than two pushes through the channel

	currentProcesses := 0
	if head.Right != nil{
		currentProcesses++
		go DFSHelper(input, append(stack, head.Right), c)
	}

	if head.Left != nil{
		currentProcesses++
		go DFSHelper(input, append(stack, head.Left), c)
	}

	for len(c) != currentProcesses{
		select{
		case c <- false:
			fmt.Println("Found one of the children to be false")
		case c <- true:
			fmt.Println("Found one of the children to be false")
			return nil
		default:
			continue
		}
	}

	return errors.New("Node not found")

}

func DFSHelper(input int, stack []*Node, done chan bool){
	if stack[len(stack)-1].Priority == input{
		done <- true
		return
	}else{
		if stack[len(stack)-1].Left != nil{
			stack = append(stack,stack[len(stack)-1].Left)
			DFSHelper(input, stack,done)
		}
	}
}
