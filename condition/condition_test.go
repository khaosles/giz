package condition

import (
	"fmt"
	"testing"
	"time"
)

/*
   @File: condition_test.go
   @Author: khaosles
   @Time: 2023/8/13 09:30
   @Desc:
*/

func TestNor(t *testing.T) {

	s := time.Now()
	a := 1
	b := 2
	for i := 0; i < 10000000; i++ {
		//Nor(a, b)
		_ = !(a != 0) || (b != 0)
	}
	fmt.Printf("expend: %f s\n", time.Now().Sub(s).Seconds())

}
