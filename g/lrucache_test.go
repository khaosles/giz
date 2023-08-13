package g

import (
	"fmt"
	"testing"
)

/*
   @File: lrucache_test.go
   @Author: khaosles
   @Time: 2023/8/12 23:52
   @Desc:
*/

func TestLRUCache(t *testing.T) {
	t.Parallel()

	cache := NewLRUCache[int, int](3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	v, ok := cache.Get(1)
	fmt.Println(v, ok)
	v, ok = cache.Get(2)
	fmt.Println(v, ok)

	ok = cache.Delete(2)
	fmt.Println(ok)

	_, ok = cache.Get(2)
	fmt.Println(ok)

}
