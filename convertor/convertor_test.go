package convertor

import (
	"encoding/binary"
	"fmt"
	"testing"
	"time"
)

/*
   @File: convertor_test.go
   @Author: khaosles
   @Time: 2023/8/10 23:28
   @Desc:
*/

type b struct {
	A string
	C int
	E string
}

func TestToBytes(t *testing.T) {
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
		//toByte(111)
		var b b
		CopyProperties(&b, a)
		//ToBytes(111)
	}
	fmt.Printf("run: %f\n", time.Now().Sub(st).Seconds())
}

func toByte(val int16) []byte {
	// 创建一个长度为 2 的 byte 数组
	byteArray := make([]byte, 2)
	// 将 int16 转换为 byte 数组
	binary.LittleEndian.PutUint16(byteArray, uint16(val))
	return byteArray
}
