package basicsorting

func BubbleSort(array []int) []int {
	for i := len(array) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}

	return array
}

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		temp := array[i]
		array[i] = array[min]
		array[min] = temp

	}

	return array
}

func InsertionSort(array []int) []int {
	//          *3
	// 3, 4, 5, 2, 1, 9, 5, 6, 7, 8, 7, 5, 4, -7, 3, 232, 3, 45, 56
	//       *2
	// 3, 4, 2, 5, 1, 9
	//    *
	// 3, 2, 4, 5, 1, 9

	for i := 1; i < len(array); i++ {
		tempIndex := i
		for j := i - 1; j >= 0; j-- {
			if array[tempIndex] < array[j] {
				temp := array[j]
				array[j] = array[tempIndex]
				array[tempIndex] = temp

				tempIndex--
			} else {
				break
			}
		}
	}
	return array
}
