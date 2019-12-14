package main

import(
	"fmt"
)

func main(){
	fmt.Println("Compiled")
	head := Node{
		Priority:3,
		Right: nil,
		Left: nil,
	}
	head.Insert(1)
	head.Insert(0)
	head.Insert(2)
	head.Insert(5)
	head.Insert(4)
	head.Insert(6)
	head.Print()
	fmt.Println()
	head.Remove(1)
	head.Remove(0)
	head.Print()
}
