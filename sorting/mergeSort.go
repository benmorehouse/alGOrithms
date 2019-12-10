package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Compiled Successfully")
	rand.Seed(time.Now().UTC().UnixNano())
	var Array []int
	for i:=0;i<10;i++{
		Array = append(Array,rand.Intn(100))
		fmt.Print(Array[i]," ")
	}
	fmt.Println()
	MergeSort(Array)
	for i:=0;i<10;i++{
		fmt.Print(Array[i]," ")
	}
}

func MergeSort(Array []int){
	// here we break all of the items down first
	MergeSortHelper(Array,0,len(Array)-1)
}

func MergeSortHelper(Array []int, LeftIndex int, RightIndex int){
	m := (LeftIndex + RightIndex)/2
	if LeftIndex < RightIndex{
		MergeSortHelper(Array,LeftIndex,m)
		MergeSortHelper(Array,m+1,RightIndex)
	}
	fmt.Println(LeftIndex, RightIndex)
	Merge(Array[LeftIndex:m], Array[m:RightIndex])
}

func Merge(LeftArray, RightArray []int)([]int){
	fmt.Println("Left array:",LeftArray)
	fmt.Println("Right array:",RightArray)
	var Array []int
	L := 0
	R := 0
	for L != len(LeftArray) && R != len(RightArray){
		if LeftArray[L] < RightArray[R]{
			Array = append(Array,LeftArray[L])
			L++
		}else{
			Array = append(Array,RightArray[R])
		}
	}

	if L == len(LeftArray){
		for R != len(RightArray){
			Array = append(Array,RightArray[R])
			R++
		}
	}else{
		for L != len(LeftArray){
			Array = append(Array,LeftArray[R])
			L++
		}
	}
	return Array
}

