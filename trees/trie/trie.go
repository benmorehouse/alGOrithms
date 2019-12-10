package main

import(
	"errors"
	"fmt"
)

type Node struct{
	Letter rune
	Letters map[rune]*Node
}

type Trie struct{
	Root *Node
}

func (this *Trie) Insert(_i interface{})error{
	input, ok := _i.(string)
	if ok == false{
		return errors.New("Entered interface is not of the correct type")
	}

	if this.Root == nil{
		newRootNode := Node{
			Letter: '0',
			Letters: make(map[rune]*Node),
		}
		this.Root = &newRootNode
	}
	currentNode := this.Root
	i := 0
	for i < len(input){
		// go through, see if it returns nil or not
		if currentNode.Letters[rune(input[i])] == nil{
			// then make that node and move to it
			newNode := Node{
				Letter: rune(input[i]),
				Letters: make(map[rune]*Node),
			}
			currentNode.Letters[rune(input[i])] = &newNode
			currentNode = &newNode
			i++
		}else{
			currentNode = currentNode.Letters[rune(input[i])]
			i++
		}
	}
	return nil
}

func (this *Trie) Search(_i interface{})bool{
	input, ok := _i.(string)
	if ok == false{
		return false
	}

	if this.Root == nil{
		return false
	}
	currentNode := this.Root
	for i:=0;i<len(input);i++{
		currentNode = currentNode.Letters[rune(input[i])]
		if currentNode == nil{
			return false
		}
	}
	return true
}

func (this *Trie) Display(){
	// this should probably be a backtracking problem
	this.DisplayHelper("",this.Root)
}

func (this *Trie) DisplayHelper(currentString string, currentNode *Node){
	end := true
	for i:='a';i<'z'+1;i++{
		if currentNode.Letters[rune(i)] != nil{
			end = false
			this.DisplayHelper(currentString + string(i),currentNode.Letters[rune(i)])
		}
	}
	if end{
		// then print and backtrack
		fmt.Println(currentString)
	}
	return
}
