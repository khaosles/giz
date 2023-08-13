package jsontuil

import (
	"github.com/bytedance/sonic"
)

/*
   @File: json.go
   @Author: khaosles
   @Time: 2023/6/16 19:28
   @Desc:
*/

func ParseObject(jsonStr string) (*JsonObject, error) {
	value := new(Value)

	err := sonic.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}

	return value.JsonObject(), nil
}

func ParseArray(jsonStr string) (*JsonArray, error) {
	value := new(Value)

	err := sonic.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}
	return value.JsonArray(), nil
}
