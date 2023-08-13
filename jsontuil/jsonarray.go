package jsontuil

import (
	"sync"

	"github.com/bytedance/sonic"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: jsonarray.go
   @Author: khaosles
   @Time: 2023/6/16 19:28
   @Desc:
*/

type JsonArray struct {
	s     []*Value
	mutex sync.Mutex
}

func (j *JsonArray) Get(index int) (*Value, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if index >= len(j.s) {
		return nil, IndexOutOfRangeError
	}
	return j.s[index], nil
}

func (j *JsonArray) Value() []*Value {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	return j.s
}

func (j *JsonArray) GetJsonObject(index int) *JsonObject {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		panic(err)
	}
	return val.JsonObject()
}

func (j *JsonArray) GetJsonArray(index int) *JsonArray {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		panic(err)
	}
	return val.JsonArray()
}

func (j *JsonArray) GetString(index int) (string, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return "", err
	}
	return val.String()
}

func (j *JsonArray) GetInt(index int) (int, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return 0, err
	}
	return val.Int()
}

func (j *JsonArray) GetNullInt(index int) (*Integer, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullInt()
}

func (j *JsonArray) GetInt64(index int) (int64, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return 0, err
	}
	return val.Int64()
}

func (j *JsonArray) GetNullLong(index int) (*Long, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullLong()
}

func (j *JsonArray) GetFloat64(index int) (float64, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return 0, err
	}
	return val.Float64()
}

func (j *JsonArray) GetNullFloat(index int) (*Float, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullFloat()
}

func (j *JsonArray) GetBoolean(index int) (bool, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return false, err
	}
	return val.Boolean()
}

func (j *JsonArray) GetNullBoolean(index int) (*Boolean, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullBoolean()
}

func (j *JsonArray) String() string {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if j.s == nil {
		return ""
	}
	data, err := sonic.Marshal(j.s)
	if err != nil {
		glog.Error(err)
		return ""
	}
	return string(data)
}
