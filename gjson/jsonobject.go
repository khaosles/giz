package gjson

import (
	"sort"
	"sync"

	"github.com/bytedance/sonic"
)

/*
   @File: jsonobject.go
   @Author: khaosles
   @Time: 2023/6/16 19:28
   @Desc:
*/

type JsonObject struct {
	m     map[string]*Value
	mutex sync.RWMutex
}

func NewJsonObject() *JsonObject {
	return &JsonObject{m: map[string]*Value{}}
}

func (j *JsonObject) HasKey(key string) bool {
	if j.m == nil {
		return false
	}
	_, ok := j.m[key]
	return ok
}

func (j *JsonObject) Get(key string) (*Value, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if !j.HasKey(key) {
		return nil, KeyNotFoundError{key}
	}
	return j.m[key], nil
}

func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonObject()
}

func (j *JsonObject) GetJsonArray(key string) *JsonArray {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonArray()
}

func (j *JsonObject) GetString(key string) (string, error) {
	val, err := j.Get(key)
	if err != nil {
		return "", err
	}
	return val.String()
}

func (j *JsonObject) GetInt(key string) (int, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int()
}

func (j *JsonObject) GetNullInt(key string) (*Integer, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullInt()
}

func (j *JsonObject) GetInt64(key string) (int64, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int64()
}

func (j *JsonObject) GetNullLong(key string) (*Long, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullLong()
}

func (j *JsonObject) GetFloat64(key string) (float64, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Float64()
}

func (j *JsonObject) GetNullFloat(key string) (*Float, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullFloat()
}

func (j *JsonObject) GetBoolean(key string) (bool, error) {
	val, err := j.Get(key)
	if err != nil {
		return false, err
	}
	return val.Boolean()
}

func (j *JsonObject) GetNullBoolean(key string) (*Boolean, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullBoolean()
}

func (j *JsonObject) Put(key string, val interface{}) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	j.m[key] = &Value{val}
}

func (j *JsonObject) String() string {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if j.m == nil {
		return ""
	}
	data, err := sonic.Marshal(j.m)
	if err != nil {
		return ""
	}
	return string(data)
}

func (j *JsonObject) Values() map[string]*Value {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	return j.m
}

func (j *JsonObject) Sort() *JsonObject {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	// 提取 map 中的键到切片
	keys := make([]string, 0, len(j.m))
	for key := range j.m {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Strings(keys)
	newM := make(map[string]*Value)
	// 根据排序后的键遍历输出值
	for _, key := range keys {
		newM[key] = j.m[key]
	}
	j.m = newM
	return j
}

// MarshalJSON implements the json.Marshaler interface for JsonObject.
func (j *JsonObject) MarshalJSON() ([]byte, error) {
	defer j.mutex.RUnlock()
	j.mutex.RLock()
	return sonic.Marshal(j.m)
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsonObject.
func (j *JsonObject) UnmarshalJSON(data []byte) error {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	j.m = make(map[string]*Value) // Reset the map
	if err := sonic.Unmarshal(data, &j.m); err != nil {
		return err
	}
	return nil
}
