package mergesort

import (
	"fmt"
	"math"
)

func MergeSort(array []int) []int {
	fmt.Println("Input array: \n", array)
	if len(array) == 1 {
		return array
	}

	middle := int(math.Floor(float64(len(array)) / 2))
	fmt.Println("Middle: ", middle)
	left := array[:middle]
	right := array[middle:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(arr1, arr2 []int) []int {
	combined := []int{}
	i := 0
	j := 0

	// Doing util it is empty
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			combined = append(combined, arr1[i])
			i++
		} else {
			combined = append(combined, arr2[j])
			j++
		}
	}

	// Doing what is rested in array
	for i < len(arr1) {
		combined = append(combined, arr1[i])
		i++
	}

	for j < len(arr2) {
		combined = append(combined, arr2[j])
		j++
	}

	return combined
}
