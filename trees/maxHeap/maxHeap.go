package maxheap

import(
	"errors" // Insert type assertions
	"fmt" //Display
	"math" // max
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

	this.Array = append(this.Array,&newNode)
	this.TrickleUp(len(this.Array)-1)// we now trickle up the heap
	return nil
}

func (this *Heap) Peek()(*Node){
	if len(this.Array) == 0{
		return nil
	}

	return this.Array[0]
}

func (this *Heap) Display(){
	for _,val := range this.Array{
		fmt.Println(val.Priority)
	}
}

func (this *Heap) TrickleUp(index int){
	if index < 0 || index >= len(this.Array){
		return
	}

	parentIndex := (index-1)/2

	for this.Array[index].Priority > this.Array[parentIndex].Priority{
		this.Array[index],this.Array[parentIndex] = this.Array[parentIndex],this.Array[index]
		index, parentIndex = parentIndex, ((parentIndex-1)/2)
	}// move things to the top that need to be moved
}

func (this *Heap) Pop(){
	// switch the last and first array index, pop last array, then trickle down array[0]
	if len(this.Array) == 0{
		return
	}else if len(this.Array) == 1{
		this.Array = this.Array[:len(this.Array)-1]
		return
	}
	this.Array[0], this.Array[len(this.Array)-1] = this.Array[len(this.Array)-1],this.Array[0]
	this.Array = this.Array[:len(this.Array)-1]
	this.TrickleDown(0)
}

func (this *Heap) TrickleDown(index int){
	if index < 0 || index >= len(this.Array){
		return
	}else if len(this.Array) == 1{
		return
	}else if len(this.Array) == 2{
		if this.Array[0].Priority > this.Array[1].Priority{
			this.Array[0], this.Array[1] = this.Array[1], this.Array[0]
		}
		return
	}

	leftChild := (2*index)+1
	rightChild := (2*index)+2
	done := false
	for !done{
		if leftChild < len(this.Array) && rightChild < len(this.Array){
			largestChild := math.Max(float64(this.Array[leftChild].Priority),float64(this.Array[rightChild].Priority))
			if largestChild < float64(this.Array[index].Priority){
				if largestChild == float64(this.Array[leftChild].Priority){
					this.Array[index], this.Array[leftChild] = this.Array[leftChild],this.Array[index]
					index = leftChild
				}else{
					this.Array[index], this.Array[rightChild] = this.Array[rightChild],this.Array[index]
					index = rightChild
				}
			}else{
				break
			}
			leftChild = (2*index)+1
			rightChild = (2*index)+2
		}else if leftChild < len(this.Array){
			if this.Array[leftChild].Priority > this.Array[index].Priority{
				this.Array[index], this.Array[leftChild] = this.Array[leftChild],this.Array[index]
				break
			}
		}else{
			break
		}
	}
}

