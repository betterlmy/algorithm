package search

func BinarySearchRecursion(ints []int, target int) int {
	start, end := 0, len(ints)-1
	if (end - start) == 0 {
		return -1
	}
	return searchRecursion(ints, target, start, end)

}

func searchRecursion(ints []int, target int, startIndex int, endIndex int) int {
	if endIndex == startIndex {
		if ints[endIndex] != target {
			return -1
		} else {
			return endIndex
		}
	}
	mid := (endIndex + startIndex) / 2
	tmp := ints[mid]

	if tmp == target {
		return mid
	} else if tmp > target {
		return searchRecursion(ints, target, startIndex, mid)
	} else {
		return searchRecursion(ints, target, mid+1, endIndex)
	}

}
