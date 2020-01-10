package hash_count

// Hash will receive a hash and return the count of each string in it
func Hash(array []string) map[string]int {

	countedStringMap := make(map[string]int)
	for _, value := range array {
		countedStringMap[value]++
	}
	return countedStringMap
}
