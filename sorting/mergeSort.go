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
	for i:=0;i<3;i++{
		Array = append(Array,rand.Intn(100))
		fmt.Print(Array[i]," ")
	}
	fmt.Println()
	MergeSort(Array)
	for i:=0;i<3;i++{
		fmt.Print(Array[i]," ")
	}
}

func MergeSort(Array []int){
	// here we break all of the items down first
	MergeSortHelper(Array,0,len(Array)-1)
}

func MergeSortHelper(Array []int, LeftIndex int, RightIndex int){
	m := (LeftIndex + RightIndex)/2
	fmt.Println("m is:",m)
	fmt.Println("LeftIndex:",LeftIndex)
	fmt.Println("RightIndex:",RightIndex)
	if LeftIndex < RightIndex{
		MergeSortHelper(Array,LeftIndex,m)
		MergeSortHelper(Array,m+1,RightIndex)
	}
	Merge(Array[LeftIndex:m], Array[m:RightIndex])
}

func Merge(LeftArray, RightArray []int, LeftIndex, RightIndex int)([]int){
	fmt.Println("Left array:",LeftArray)
	fmt.Println("Right array:",RightArray)
	var Array []int
	L := 0
	R := 0

	if len(LeftArray) == 0 || len(RightArray) == 0{
		fmt.Println("Hit first if statement")
		if len(LeftArray) == 0{
			return RightArray
		}
		return LeftArray
	}
	fmt.Println("Hit for loop")

	for L < len(LeftArray) && R < len(RightArray){
		if LeftArray[L] < RightArray[R]{
			Array = append(Array,LeftArray[L])
			L++
		}else{
			Array = append(Array,RightArray[R])
			R++
		}
	}

	if L == len(LeftArray){
		fmt.Println("Hit second if statement")
		for R != len(RightArray)-1{
			Array = append(Array,RightArray[R])
			R++
		}
	}else{
		fmt.Println("Hit third if statement")
		for L != len(LeftArray)-1{
			Array = append(Array,LeftArray[R])
			L++
		}
	}
	fmt.Println("Array at end of merge function:",Array)
	return Array
}

