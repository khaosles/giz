package algorithm

import "github.com/khaosles/giz/constraints"

/*
   @File: search.go
   @Author: khaosles
   @Time: 2023/8/13 13:54
   @Desc:
*/

// LinearSearch return the index of target in slice base on equal function.
// If not found return -1
func LinearSearch[T any](slice []T, target T, equal func(a, b T) bool) int {
	for i, v := range slice {
		if equal(v, target) {
			return i
		}
	}
	return -1
}

// BinarySearch return the index of target within a sorted slice, use binary search (recursive call itself).
// If not found return -1.
func BinarySearch[T any](sortedSlice []T, target T, comparator constraints.Comparator[T]) int {
	return binarySearch[T](sortedSlice, target, 0, len(sortedSlice)-1, comparator)
}

func binarySearch[T any](sortedSlice []T, target T, low, hight int, comparator constraints.Comparator[T]) int {
	if hight < low || len(sortedSlice) == 0 {
		return -1
	}
	mid := (hight + low) >> 1
	if comparator.Compare(sortedSlice[mid], target) > 0 {
		return binarySearch(sortedSlice, target, low, mid-1, comparator)
	} else if comparator.Compare(sortedSlice[mid], target) < 0 {
		return binarySearch(sortedSlice, target, mid+1, hight, comparator)
	}
	return mid
}

// BinaryIterativeSearch return the index of target within a sorted slice, use binary search (no recursive).
// If not found return -1.
func BinaryIterativeSearch[T any](sortedSlice []T, target T, comparator constraints.Comparator[T]) int {
	return binaryIterativeSearch[T](sortedSlice, target, 0, len(sortedSlice)-1, comparator)
}

func binaryIterativeSearch[T any](sortedSlice []T, target T, low, hight int, comparator constraints.Comparator[T]) int {
	var mid int
	for low <= hight {
		mid = (low + hight) >> 1
		if comparator.Compare(sortedSlice[mid], target) > 0 {
			hight = mid - 1
		} else if comparator.Compare(sortedSlice[mid], target) < 0 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
