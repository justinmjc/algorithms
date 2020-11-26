package skiplist

import (
	"math/rand"
	"sync"
	"time"
)
//使用go语言实现跳跃表的增、删、改、查操作
//考虑到业务上的使用场景，该实现将支持协程安全。

//节点结构体
type SkipListNode struct {
	key int
	data interface{}
	next []*SkipListNode //next为节点指针切片 比如图中的"HEAD"节点的next应该为[&5, &3, &1, &1]，“3”节点的next应该为[&5, &5, &4]
}

//skiplist 结构体

type SkipList struct {
	head *SkipListNode
	tail *SkipListNode
	length int
	level int
	mut *sync.RWMutex //互斥锁，用于保证协程安全
	rand *rand.Rand //随机数生成器，用于生成随机层数，随机生成的层数要满足P=0.5的几何分布，大致可以理解为：掷硬币连续出现正面的次数，就是我们要的层数。
}

func NewSkipList(level int)*SkipList  {
	list := &SkipList{}
	if level<=0{
		level =32
	}
	list.level = level

	list.head = &SkipListNode{
		next: make([]*SkipListNode,level,level),
	}

	list.tail = &SkipListNode{}

	list.mut = &sync.RWMutex{}

	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))


	return list
}