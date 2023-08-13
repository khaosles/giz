package algorithm

import (
	"fmt"
	"testing"
)

/*
   @File: search_test.go
   @Author: khaosles
   @Time: 2023/8/13 16:06
   @Desc:
*/

func TestLinearSearch(t *testing.T) {
	numbers := []int{3, 4, 5, 3, 2, 1}

	equalFunc := func(a, b int) bool {
		return a == b
	}

	result1 := LinearSearch(numbers, 3, equalFunc)
	result2 := LinearSearch(numbers, 6, equalFunc)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 0
	// -1
}

func TestBinarySearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result1 := BinarySearch[int](numbers, 5, NumericComparator[int]{})
	result2 := BinarySearch[int](numbers, 9, NumericComparator[int]{})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 4
	// -1
}

func TestBinaryIterativeSearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result1 := BinaryIterativeSearch[int](numbers, 5, NumericComparator[int]{})
	result2 := BinaryIterativeSearch[int](numbers, 9, NumericComparator[int]{})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 4
	// -1
}
