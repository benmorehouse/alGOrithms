package maxheap_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "." //import for access to rest of built packages
)

func TestMaxHeap(t *testing.T) {
	RegisterFailHandler(Fail) // registers the fail handler from ginkgo
	RunSpecs(t, "MaxHeap tests ") // hands over control to the ginkgo testing framework
}

func HeapHolds(h *Heap){
	holds := true
	for i := 0; i < len(h.Array); i++{
		left := 2*i+1
		right := 2*i+2
		if left >= len(h.Array){
			continue
		}else{
			Expect(h.Array[i].Priority >= h.Array[left].Priority).Should(BeTrue())
			if h.Array[i].Priority < h.Array[left].Priority{
				holds = false
				break
			}
		}

		if right >= len(h.Array){
			continue
		}else{
			Expect(h.Array[i].Priority >= h.Array[right].Priority).Should(BeTrue())
			if h.Array[i].Priority < h.Array[left].Priority{
				holds = false
				break
			}
		}
	}
	Expect(holds).Should(BeTrue())
}

var _ = Describe("Maxheap initialization", func(){ //describe is the exact same functionality as the Context, but semantically is most common to see this.
	h, err := NewHeap()
	Expect(err).Should(BeNil())
	if err != nil{
		return
	}

	Context("init", func(){// context allows you to organize your specs(ie things like It
		It("heap should have zero elements", func(){
			Expect(h.Array).Should(HaveLen(0))
		}) // it code contains all your actual tests and assertions of what is happening
	})

	h, err = NewHeap(1,2,3,4)
	Expect(err).Should(BeNil())
	if err != nil{
		return
	}

	Context("Check init function retains heap property",func(){
		Specify("Heap should have exact four elements",func(){
			Expect(h.Array).Should(HaveLen(4))
		})

		It("Heap property still holds",func(){
			HeapHolds(h)
		})
	})
}

var _ = Describe("Maxheap operations with empty heap", func(){
	Context("popping empty heap", func(){
		Specify("heap should have no elements in it", func(){
			Expect(0).Should(BeZero())
		})
		It("function should return right away", func(){})
	})

	Context("pushing empty heap", func(){
		Specify("shouldnt have any nodes",func(){
			Expect(0).Should(BeZero())
		})
		It("should increase size of internal array by one",func(){})
		Specify("root node should now be whatever was just added",func(){})
	})
})

/*
	Context("popping non-empty heap", func(){
		It("heap should have at least one element",func(){})
		It("heap should have one less element than it did earlier",func(){})
		It("heap should retain heap properties", func(){

		})
	})

	Context("pushing non-empty heap", func(){
		It("should have some nodes",func(){})
		It("should increase size of internal array by one",func(){})
		It("heap should retain heap properties", func(){

		})
	})

*/
