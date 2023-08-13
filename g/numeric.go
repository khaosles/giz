package g

/*
   @File: numeric.go
   @Author: khaosles
   @Time: 2023/8/12 15:57
   @Desc:
*/

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
