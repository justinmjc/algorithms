# 数组
## 数组结构特性
1. 查询很快，时间复杂度O(1)
2. 插入，删除，插入后的所有元素要后移，时间复杂度O(n),为了弥补这一不足，出现了链表

## 思考
### 1. 数组元素的删除怎么做？
```
删除 第i个元素
arr = append(arr[:i-1],arr[i:]...)
//删除第2个元素
array := []int{1,3,5}
array = append(array[:1],array[2:]... )

```
### 2. 数组元素的插入？
```
在末尾插入
array = append(array[:1],6 )
//在第i个位置插入
array := []int{1,3,5}
array = append(array[:1],array[2:]... )

```