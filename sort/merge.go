package sort

func MergeSort(ints []int) {
	mergeSort(ints, 0, len(ints)-1)
}

func mergeSort(ints []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (end + start) / 2
	mergeSort(ints, start, mid)
	mergeSort(ints, mid+1, end)
	merge(ints, start, mid, end)
}

func merge(ints []int, start int, mid int, end int) {
	tmp := make([]int, 0)
	// 临时存储结果的变量
	i, j := start, mid+1

	for i <= mid && j <= end {
		// 分别迭代两个子串,向tmp中写入有顺序的元素
		if ints[i] > ints[j] {
			tmp = append(tmp, ints[j])
			j++
			if j > end {
				tmp = append(tmp, ints[i:mid+1]...)
				break
			}
		} else {
			tmp = append(tmp, ints[i])
			i++
			if i > mid {
				tmp = append(tmp, ints[j:end+1]...)
				break
			}
		}
	}
	oriIndex, tmlIndex := start, 0
	for oriIndex <= end {
		// 将tmp中的元素更新到原数组中
		ints[oriIndex] = tmp[tmlIndex]
		oriIndex++
		tmlIndex++
	}

}
