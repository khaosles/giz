package stream

import (
	"fmt"
	"testing"
)

/*
   @File: stream_test.go
   @Author: khaosles
   @Time: 2023/8/13 13:46
   @Desc:
*/

func TestOf(t *testing.T) {
	s := Of(1, 2, 3)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func TestFromSlice(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func TestFromChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 1; i < 4; i++ {
			ch <- i
		}
		close(ch)
	}()

	s := FromChannel(ch)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func TestFromRange(t *testing.T) {
	s := FromRange(1, 5, 1)

	data := s.ToSlice()
	fmt.Println(data)

	// Output:
	// [1 2 3 4 5]
}

func TestGenerate(t *testing.T) {
	n := 0
	max := 4

	generator := func() func() (int, bool) {
		return func() (int, bool) {
			n++
			return n, n < max
		}
	}

	s := Generate(generator)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func TestConcat(t *testing.T) {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6})

	s := Concat(s1, s2)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3 4 5 6]
}

func TestStream_Distinct(t *testing.T) {
	original := FromSlice([]int{1, 2, 2, 3, 3, 3})
	distinct := original.Distinct()

	data1 := original.ToSlice()
	data2 := distinct.ToSlice()

	fmt.Println(data1)
	fmt.Println(data2)

	// Output:
	// [1 2 2 3 3 3]
	// [1 2 3]
}

func TestStream_Filter(t *testing.T) {
	original := FromSlice([]int{1, 2, 3, 4, 5})

	isEven := func(n int) bool {
		return n%2 == 0
	}

	even := original.Filter(isEven)

	fmt.Println(even.ToSlice())

	// Output:
	// [2 4]
}

func TestStream_Map(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	addOne := func(n int) int {
		return n + 1
	}

	increament := original.Map(addOne)

	fmt.Println(increament.ToSlice())

	// Output:
	// [2 3 4]
}

func TestStream_Peek(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	data := []string{}
	peekStream := original.Peek(func(n int) {
		data = append(data, fmt.Sprint("value", n))
	})

	fmt.Println(original.ToSlice())
	fmt.Println(peekStream.ToSlice())
	fmt.Println(data)

	// Output:
	// [1 2 3]
	// [1 2 3]
	// [value1 value2 value3]
}

func TestStream_Skip(t *testing.T) {
	original := FromSlice([]int{1, 2, 3, 4})

	s1 := original.Skip(-1)
	s2 := original.Skip(0)
	s3 := original.Skip(1)
	s4 := original.Skip(5)

	fmt.Println(s1.ToSlice())
	fmt.Println(s2.ToSlice())
	fmt.Println(s3.ToSlice())
	fmt.Println(s4.ToSlice())

	// Output:
	// [1 2 3 4]
	// [1 2 3 4]
	// [2 3 4]
	// []
}

func TestStream_Limit(t *testing.T) {
	original := FromSlice([]int{1, 2, 3, 4})

	s1 := original.Limit(-1)
	s2 := original.Limit(0)
	s3 := original.Limit(1)
	s4 := original.Limit(5)

	fmt.Println(s1.ToSlice())
	fmt.Println(s2.ToSlice())
	fmt.Println(s3.ToSlice())
	fmt.Println(s4.ToSlice())

	// Output:
	// []
	// []
	// [1]
	// [1 2 3 4]
}

func TestStream_AllMatch(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	result1 := original.AllMatch(func(item int) bool {
		return item > 0
	})

	result2 := original.AllMatch(func(item int) bool {
		return item > 1
	})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func TestStream_AnyMatch(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	result1 := original.AnyMatch(func(item int) bool {
		return item > 1
	})

	result2 := original.AnyMatch(func(item int) bool {
		return item > 3
	})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func TestStream_NoneMatch(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	result1 := original.NoneMatch(func(item int) bool {
		return item > 3
	})

	result2 := original.NoneMatch(func(item int) bool {
		return item > 1
	})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func TestStream_ForEach(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	result := 0
	original.ForEach(func(item int) {
		result += item
	})

	fmt.Println(result)

	// Output:
	// 6
}

func TestStream_Reduce(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	result := original.Reduce(0, func(a, b int) int {
		return a + b
	})

	fmt.Println(result)

	// Output:
	// 6
}

func TestStream_FindFirst(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	result, ok := original.FindFirst()

	fmt.Println(result)
	fmt.Println(ok)

	// Output:
	// 1
	// true
}

func TestStream_FindLast(t *testing.T) {
	original := FromSlice([]int{3, 2, 1})

	result, ok := original.FindLast()

	fmt.Println(result)
	fmt.Println(ok)

	// Output:
	// 1
	// true
}

func TestStream_Reverse(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	reverse := original.Reverse()

	fmt.Println(reverse.ToSlice())

	// Output:
	// [3 2 1]
}

func TestStream_Range(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})

	s1 := original.Range(0, 0)
	s2 := original.Range(0, 1)
	s3 := original.Range(0, 3)
	s4 := original.Range(1, 2)

	fmt.Println(s1.ToSlice())
	fmt.Println(s2.ToSlice())
	fmt.Println(s3.ToSlice())
	fmt.Println(s4.ToSlice())

	// Output:
	// []
	// [1]
	// [1 2 3]
	// [2]
}

func TestStream_Sorted(t *testing.T) {
	original := FromSlice([]int{4, 2, 1, 3})

	sorted := original.Sorted(func(a, b int) bool { return a < b })

	fmt.Println(original.ToSlice())
	fmt.Println(sorted.ToSlice())

	// Output:
	// [4 2 1 3]
	// [1 2 3 4]
}

func TestStream_Max(t *testing.T) {
	original := FromSlice([]int{4, 2, 1, 3})

	max, ok := original.Max(func(a, b int) bool { return a > b })

	fmt.Println(max)
	fmt.Println(ok)

	// Output:
	// 4
	// true
}

func TestStream_Min(t *testing.T) {
	original := FromSlice([]int{4, 2, 1, 3})

	min, ok := original.Min(func(a, b int) bool { return a < b })

	fmt.Println(min)
	fmt.Println(ok)

	// Output:
	// 1
	// true
}

func TestStream_Count(t *testing.T) {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{})

	fmt.Println(s1.Count())
	fmt.Println(s2.Count())

	// Output:
	// 3
	// 0
}
