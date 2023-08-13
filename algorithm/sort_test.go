package algorithm

import (
	"fmt"
	"testing"
	"time"
)

/*
   @File: sort_test.go
   @Author: khaosles
   @Time: 2023/8/12 13:42
   @Desc:
*/

func TestBubbleSort(t *testing.T) {
	a := []int{1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	BubbleSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestQuickSort(t *testing.T) {
	a := []int{1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	QuickSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestInsertionSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	InsertionSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestInsertionSortHalf(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	InsertionSortHalf[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestSelectionSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	SelectionSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestCountSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	CountSort(a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestMergeSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	MergeSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestMergeSortLoop(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	MergeSortLoop[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestShellSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	ShellSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestHeapSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	HeapSort[int](a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestRadixSort(t *testing.T) {
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	RadixSort(a, NumericComparator[int]{})
	fmt.Printf("%+v\n", a)
}

func TestName(t *testing.T) {
	s := time.Now()
	a := []int{2, 0, 1, 4, 5, 6, 78, 2, 21, 57, 9, 82, 58, 177, 10}
	for i := 0; i < 1000000; i++ {
		InsertionSortHalf[int](a, NumericComparator[int]{})
	}
	fmt.Printf("expend: %f s\n", time.Now().Sub(s).Seconds())
}
