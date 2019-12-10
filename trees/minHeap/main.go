package main

import(
	"fmt"
)

func main(){
	fmt.Println("Compiled")
	var min Heap
	min.Insert(12)
	min.Insert(2)
	min.Insert(1)
	min.Insert(11)
	min.Insert(0)
	min.Insert(3)
	min.Insert(123)
	min.Insert(9)
	min.Display()
}
