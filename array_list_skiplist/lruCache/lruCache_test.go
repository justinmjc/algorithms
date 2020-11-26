package lruCache

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T)  {
	cache := NewLruCache(512)

	key := "k1"
	cache.Set(key,1)
	fmt.Printf("cache 元素个数：%d, 占用内存 %d 字节\n\n", cache.Len(), cache.UseBytes())

	val := cache.Get(key)
	fmt.Println(val)
	cache.Deluseless()
	fmt.Printf("cache 元素个数：%d, 占用内存 %d 字节\n\n", cache.Len(), cache.UseBytes())
}
