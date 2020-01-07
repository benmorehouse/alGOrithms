package maxheap_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "." //import otherwise we will not be able to get access to heap file properties
)

func TestMaxHeap(t *testing.T) {
	RegisterFailHandler(Fail) // registers the fail handler from ginkgo
	RunSpecs(t, "MaxHeap Suite") // hands over control to the ginkgo testing framework
}

var _ = Describe("Maxheap initialization", func(){ //describe is the exact same functionality as the Context, but semantically is most common to see this.
	h, err := NewHeap()
	Expect(err).Should(BeNil())
	if err != nil{
		return
	}

	Context("init", func(){// context allows you to organize your specs(ie things like It
		It("heap should have zero elements", func(){
			Expect(h.Array).Should(HaveLen(1))
		}) // it code contains all your actual tests and assertions of what is happening
	})

	h, err = NewHeap(1,2,3,4)
	Expect(err).Should(BeNil())
	if err != nil{
		return
	}

	Context("Check init function retains heap property",func(){
		It("Heap should have exact four elements",func(){
			Expect(h.Array).Should(HaveLen(1))
		})

		It("Heap property still holds",func(){
			for i := 0; i < len(*h.Array); i++{
				left := 2*i+1
				right := 2*i+2
				if left > len(*h.Array){
					continue
				}else{
					if *h.Array[i] < *h.Array[left]{
						return false
					}
				}

				if right > len(h.Array){
					continue
				}else{
					if h.Array[i] < h.Array[right]{
						return false
					}
				}
			}
			return true
		})
	})
})

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
