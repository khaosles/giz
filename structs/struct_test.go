package structs

import (
	"fmt"
	"testing"
	"time"
)

/*
   @File: struct_test.go
   @Author: khaosles
   @Time: 2023/8/11 00:38
   @Desc:
*/

type b struct {
	A string
	C int
	E string
}

func TestToMap(t *testing.T) {
	a := struct {
		A string
		B string
		C int
		D float32
	}{
		"a", "b", 1, 1.0,
	}
	st := time.Now()
	for i := 0; i < 10000000; i++ {
		var b b
		err := CopyProperties(&b, a)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("run: %f\n", time.Now().Sub(st).Seconds())

}
