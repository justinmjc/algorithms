package lruCache

import (
	"container/list"
	"fmt"
	"runtime"
)

//定义Cache接口
type Cache interface {
	Set(key string,value interface{})
	Get(key string) interface{}
	Del(key string)
	//"删除访问频次最低的元素"
	Deluseless()
	Len() int
	//缓存大小
	UseBytes() int
}

//结构体 数组 切片 map 要求实现value 接口，返回占用内存的字节数
type Value interface {
	Len() int
}

//定义 key ,value 结构
type entry struct {
	key string
	value interface{}
}

//计算出元素占用内存字节数
func (e *entry)Len() int{
	return CalcLen(e.value)
}

// 计算value占用内存大小
func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value: // 结构体，数组，切片，map,要求实现 Value 接口，该接口只有1个 Len 方法，返回占用的内存字节数，如果没有实现该接口，则panic
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, int8, uint8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.value", value))
	}

	return n
}

type lru struct {
	maxBytes int
	usedBytes int
	//双链表
	ll *list.List
	// map的key是字符串，value是双链表中对应节点的指针
	cache map[string]*list.Element
}

//通过set 方法往Cache 头部增加一个元素，如有已经存在，移到头部，并更新值
func (l *lru) Set(key string, value interface{}) {
	if element,ok := l.cache[key];ok{
		//移到头部
		l.ll.MoveToFront(element)
		eval :=element.Value.(*entry)
		//重新计算内存占用
		l.usedBytes = l.usedBytes - CalcLen(eval.value) + CalcLen(value)
	}else {
		element := &entry{
			key:key,
			value: value,
		}
		e:=l.ll.PushFront(element)// 头部插入一个元素并返回该元素
		l.cache[key] = e
		l.usedBytes += element.Len()
	}

	for l.maxBytes >0 &&l.maxBytes<l.usedBytes {
		l.Deluseless()
	}
}

//获取指定元素，（有访问要将该元素移动到头部）
func (l *lru) Get(key string) interface{} {
	if e,ok := l.cache[key];ok {
		l.ll.MoveToFront(e)
		return e.Value.(*entry).value
	}
	return nil
}

func (l *lru) Del(key string) {
	if e,ok := l.cache[key];ok{
		l.removeElement(e)
	}
}

func (l *lru) Deluseless() {
	l.removeElement(l.ll.Back())
}

func (l *lru) Len() int {
	return l.ll.Len()
}

func (l *lru) UseBytes() int {
	return l.usedBytes
}

// 删除元素并更新内存占用大小
func (l *lru) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	l.ll.Remove(e)
	en := e.Value.(*entry)
	l.usedBytes -= en.Len()
	delete(l.cache, en.key)
}

func NewLruCache(maxBytes int) Cache {
	return &lru{
		maxBytes:maxBytes,
		ll: list.New(),
		cache :make(map[string]*list.Element),
	}
}