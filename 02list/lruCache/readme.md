
# lru算法 
## 核心原则
如果数据最近被访问过，那么将来被访问的概率会更高
## 实现
双链表+map 
如果某条数据被访问了，则把该条数据移动到链表头部，
队尾是最少使用的元素，内存超出限制时，淘汰队尾元素即可

1. map 用来存储键值对。这是实现缓存最简单直接的数据结构，因为它的查找记录和增加记录时间复杂度都是 O(1)

2. list.List 是go标准库提供的双链表。
通过这个双链表存放具体的值，移动任意记录到队首的时间复杂度都是 O(1)，
在队首增加记录的时间复杂度是 O(1)，删除任意一条记录的时间复杂度是 O(1)

//参考 https://blog.csdn.net/weixin_43456598/article/details/108517148