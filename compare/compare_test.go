package compare

import (
	"fmt"
	"testing"
	"time"
)

/*
   @File: compare_test.go
   @Author: khaosles
   @Time: 2023/8/10 16:06
   @Desc:
*/

func TestEqual(t *testing.T) {
	s := time.Now()
	for i := 0; i < 10000000; i++ {
		// Equal(1, 1)
		eq(1, 1)
	}
	fmt.Printf("runtime: %f\n", time.Now().Sub(s).Seconds())

}

func eq[T comparable](a, b T) bool {
	return a == b
}
