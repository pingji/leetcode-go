# Solutions

## Solution 1：
使用golang的heap, 需要实现heap接口的五个函数 Len，Less，Swap，Push，Pop

参考 [Golang 自定义最小堆](https://leetcode-cn.com/problems/top-k-frequent-words/solution/golang-zi-ding-yi-zui-xiao-dui-by-yanert-3zaz/)
### heap
heap 的 interface
```go
package heap
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
```
Push and Pop use pointer receivers because they modify the slice's length, not just its contents.

### sort
sort 的 interface
```go
package sort
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

## Solution 2：
自己实现的heap, 