// mergeSort sorts by divide and conquer
package mergeSort

// Sort an array using merge sort
func Sort(array []int) []int {
	return split(array)
}

func split(array []int) []int {

	if len(array) == 1 || len(array) == 0 {
		return array
	}

	// then at this point we need to go ahead and split the arrays up more
	left, right := []int{}, []int{}
	middleIndex := len(array) / 2
	for index, value := range array {
		if index < middleIndex {
			left = append(left, value)
			continue
		}
		right = append(right, value)
	}

	// then after splitting all of it up, merge together

	return merge(split(left), split(right))
}

func merge(leftArray, rightArray []int) []int {
	// merge the left and right together
	result := []int{}
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(leftArray) && rightIndex < len(rightArray) {
		if leftArray[leftIndex] < rightArray[rightIndex] {
			result = append(result, leftArray[leftIndex])
			leftIndex++
			continue
		}

		result = append(result, rightArray[rightIndex])
		rightIndex++
	}

	for index := rightIndex; index < len(rightArray); index++ {
		result = append(result, rightArray[rightIndex])
	}

	for index := leftIndex; index < len(leftArray); index++ {
		result = append(result, leftArray[leftIndex])
	}
	return result
}
