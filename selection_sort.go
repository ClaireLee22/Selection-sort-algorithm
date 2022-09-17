package main

import "fmt"

// Time: O(n^2) | Space: O(1)
func selectionSort(array []int) []int {
	for currentIdx := 0; currentIdx < len(array); currentIdx++ {
		smallestIdx := currentIdx
		for j := currentIdx + 1; j < len(array); j++ {
			if array[smallestIdx] > array[j] {
				smallestIdx = j
			}
		}
		swap(currentIdx, smallestIdx, array)
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 2, 5, 4, 1}
	fmt.Println(selectionSort(array))
}

/* output:
[1 2 4 5 8]
*/
