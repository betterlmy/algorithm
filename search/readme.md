# 二分查找

二分查找是在一组**有序的**数据中查找特定元素的算法.它的基本思想是将数组分成两半,然后检查中间元素是否等于目标值,如果不相等,则根据目标值与所查找到的元素之间的大小关系,如果大则继续向右查找,否则向左查找.直到找到目标值或为空.

二分查找的时间复杂度为$$O(logn)$$,比顺序查找快很多.

## 二分查找的递归实现

```go
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

```



## 二分查找的非递归实现

```go
func BinarySearch(ints []int, target int) int {
	start, end := 0, len(ints)-1

	for start <= end {
		mid := (end + start) / 2
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
```

## 二分查找的局限性

1. 二分查找依赖于有序数组,不适用于链表.
2. 数据量过小时,二分查找优势不明显.
3. 数据量太大时,一般也用不到数组,也用不到二分查找