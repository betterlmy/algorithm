package sort

func QuickSort(ints []int) []int {
	if len(ints) < 2 {
		return ints
	}

	pivot := ints[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, v := range ints[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(append(less, pivot), greater...)
}
