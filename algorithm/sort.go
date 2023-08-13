package algorithm

import (
	"math"

	"github.com/khaosles/giz/constraints"
)

/*
   @File: sort.go
   @Author: khaosles
   @Time: 2023/8/12 13:09
   @Desc:
*/

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// BubbleSort applys the bubble sort algorithm to sort the collection, will change the original collection data.
// stable
// S(n) = O(1)
// T(n)max = O(n2) T(n)mean = O(n2)
func BubbleSort[T any](arr []T, comparator constraints.Comparator[T]) {
	flag := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if comparator.Compare(arr[j], arr[j+1]) > 0 {
				swap(arr, j, j+1)
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

// QuickSort quick sorting for slice, lowIndex is 0 and highIndex is len(slice)-1.
// stable
// S(n)max = O(n)  S(n)mean = O(log2n)
// T(n)max = O(n2) T(n) = O(nlog2n)
func QuickSort[T any](arr []T, comparator constraints.Comparator[T]) {
	quickSort(arr, 0, len(arr)-1, comparator)
}

func quickSort[T any](arr []T, low, hight int, comparator constraints.Comparator[T]) {
	if low < hight {
		pivotpos := partition(arr, low, hight, comparator)
		quickSort(arr, low, pivotpos-1, comparator)
		quickSort(arr, pivotpos+1, hight, comparator)
	}
}

func partition[T any](arr []T, low, high int, comparator constraints.Comparator[T]) int {
	piovt := arr[low] // use the first element as piovt
	for low < high {
		for low < high && comparator.Compare(arr[high], piovt) >= 0 {
			high--
		}
		arr[low] = arr[high] // move the element to the left if less than poivt
		for low < high && comparator.Compare(arr[low], piovt) <= 0 {
			low++
		}
		arr[high] = arr[low] // move the element to the right if greater than poivt
	}
	arr[low] = piovt
	return low
}

// InsertionSort applys the insertion sort algorithm to sort the collection, will change the original collection data.
// stable
// S(n) = O(1)
// T(n) = O(n2)
func InsertionSort[T any](arr []T, comparator constraints.Comparator[T]) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if !(comparator.Compare(arr[j], arr[j-1]) < 0) {
				break
			}
			swap(arr, j, j-1)
		}
	}
}

func InsertionSortHalf[T any](arr []T, comparator constraints.Comparator[T]) {
	size := len(arr)
	for i := 1; i < size; i++ {
		tmpVal := arr[i]
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if comparator.Compare(arr[mid], tmpVal) > 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j := i - 1; j >= high+1; j-- {
			arr[j+1] = arr[j]
		}
		arr[high+1] = tmpVal
	}
}

// MergeSort applys the merge sort algorithm to sort the collection, will change the original collection data.
// stable
// S(n) = O(n)
// O(n) = O(nlog2n)
func MergeSort[T any](arr []T, comparator constraints.Comparator[T]) {
	mergeSort(arr, 0, len(arr)-1, comparator)
}

func MergeSortLoop[T any](arr []T, comparator constraints.Comparator[T]) {
	size := len(arr)
	for i := 1; i < size; i++ {
		low := 0
		mid := low + i - 1
		high := mid + i
		for high < size {
			merge(arr, low, mid, high, comparator)
			low = high + 1
			mid = low + i - 1
			high = mid + i
		}
		if low < size && mid < size {
			merge(arr, low, mid, size-1, comparator)
		}
	}
}

func mergeSort[T any](arr []T, low, high int, comparator constraints.Comparator[T]) {
	if low < high {
		mid := (low + high) >> 1
		mergeSort(arr, 0, mid, comparator)
		mergeSort(arr, mid+1, high, comparator)
		merge(arr, low, mid, high, comparator)
	}
}

func merge[T any](arr []T, low, mid, high int, comparator constraints.Comparator[T]) {
	i := low
	j := mid + 1
	k := 0
	tmpArr := make([]T, high-low+1)
	for i <= mid && j <= high {
		if comparator.Compare(arr[i], arr[j]) < 0 {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
		k++
	}
	for m := i; m <= mid; m++ {
		tmpArr[k] = arr[m]
		k++
	}
	for m := j; m <= high; m++ {
		tmpArr[k] = arr[m]
		k++
	}

	for m := 0; m < len(tmpArr); m++ {
		arr[low+m] = tmpArr[m]
	}
}

// RadixSort applys the radix sort algorithm to sort the collection, will change the original collection data.
// stable
// S(n) = O(r)
// T(n) = O(d(n+r))
func RadixSort(arr []int, comparator constraints.Comparator[int]) {
	size := len(arr)
	maxVal := arr[0]
	for i := 1; i < size; i++ {
		if comparator.Compare(arr[i], maxVal) > 0 {
			maxVal = arr[i]
		}
	}
	times := 1
	for maxVal/10 > 0 {
		times++
		maxVal /= 10
	}
	bucket := make([][]int, 10)
	for i := 0; i < 10; i++ {
		bucket[i] = make([]int, 0)
	}
	for i := 1; i <= times; i++ {
		for j := 0; j < size; j++ {
			ratio := arr[j] / int(math.Pow(10, float64(i-1))) % 10
			bucket[ratio] = append(bucket[ratio], arr[j])
		}
		m := 0
		for j := 0; j < 10; j++ {
			for k := 0; k < len(bucket[j]); k++ {
				arr[m] = bucket[j][k]
				m++
			}
			bucket[j] = make([]int, 0)
		}
	}
}

// SelectionSort applys the selection sort algorithm to sort the collection, will change the original collection data.
// unstable
// S(n) = O(1)
// T(n) = O(n2)
func SelectionSort[T any](arr []T, comparator constraints.Comparator[T]) {
	for i := 0; i < len(arr)-1; i++ {
		k := i
		for j := i + 1; j < len(arr); j++ {
			if comparator.Compare(arr[k], arr[j]) > 0 {
				k = j
			}
		}
		if k != i {
			swap(arr, k, i)
		}
	}
}

// CountSort applys the count sort algorithm to sort the collection, don't change the original collection data.
// unstable
// S(n) = O(k)
// T(n) = O(n+k)
func CountSort(arr []int, comparator constraints.Comparator[int]) {
	size := len(arr)
	var minVal, maxVal int
	for i := 0; i < size; i++ {
		if comparator.Compare(arr[i], minVal) < 0 {
			minVal = arr[i]
		}
		if comparator.Compare(arr[i], maxVal) > 0 {
			maxVal = arr[i]
		}
	}
	blucket := make([]int, maxVal-minVal+1)
	for i := 0; i < size; i++ {
		blucket[arr[i]-minVal]++
	}
	out := make([]int, size)
	j := 0
	for i := 0; i < len(blucket); i++ {
		for blucket[i] > 0 {
			out[j] = i + minVal
			j++
			blucket[i]--
		}
	}
}

// ShellSort applys the shell sort algorithm to sort the collection, will change the original collection data.
// ubstable
// S(n) = O(1)
// T(n) = O(n2)
func ShellSort[T any](arr []T, comparator constraints.Comparator[T]) {
	size := len(arr)
	for step := size / 2; step > 0; step /= 2 {
		for i := step; i < size; i++ {
			for j := i; j > 0 && comparator.Compare(arr[j], arr[j-step]) < 0; j -= step {
				swap(arr, j, j-step)
			}
		}
	}
}

// HeapSort applys the heap sort algorithm to sort the collection, will change the original collection data.
// unstable
// S(n) = O(1)
// T(n) = O(nlog2n)
func HeapSort[T any](arr []T, comparator constraints.Comparator[T]) {
	size := len(arr)
	for i := size>>1 - 1; i >= 0; i-- {
		heapify(arr, i, size-1, comparator)
	}
	for i := size - 1; i >= 1; i-- {
		swap(arr, 0, i)
		heapify(arr, 0, i-1, comparator)
	}
}

func heapifyLoop[T any](arr []T, i, n int, comparator constraints.Comparator[T]) {
	tmpVal := arr[i]
	child := i<<1 + 1
	for child <= n {
		if child+1 <= n && comparator.Compare(arr[child], arr[child+1]) < 0 {
			child++
		}
		if comparator.Compare(arr[child], tmpVal) <= 0 {
			break
		}
		arr[i] = arr[child]
		i = child
		child = i<<1 + 1
	}
	arr[i] = tmpVal
}

func heapify[T any](arr []T, i, n int, comparator constraints.Comparator[T]) {
	largest := i
	lson := i<<1 + 1
	rson := i<<1 + 2
	if lson < n && comparator.Compare(arr[largest], arr[lson]) < 0 {
		largest = lson
	}
	if rson < n && comparator.Compare(arr[largest], arr[rson]) < 0 {
		largest = rson
	}
	if largest != i {
		swap(arr, largest, i)
		heapify(arr, largest, n, comparator)
	}
}
