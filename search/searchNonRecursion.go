package search

func BinarySearch(ints []int, target int) int {
	start, end := 0, len(ints)-1

	for start <= end {
		mid := start + (end-start)>>1
		if ints[mid] == target {
			return mid
		} else if ints[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
