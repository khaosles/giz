package gjson

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

/*
   @File: jsonobject_test.go
   @Author: khaosles
   @Time: 2023/8/20 01:32
   @Desc:
*/

func TestJsonObject(t *testing.T) {
	s := `{"data": ["123", "456"], "msg": "123", "success": true, "code": 20000}`
	jsonObject := NewJsonObject()
	err := sonic.UnmarshalString(s, jsonObject)
	if err != nil {
		fmt.Println(err)
	}
	println(jsonObject.String())
	getString := jsonObject.GetJsonArray("data")
	fmt.Println(getString)
}
