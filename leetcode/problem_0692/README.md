# Solutions

## Solution 1：
使用golang的heap, 需要实现heap接口的五个函数 Len，Less，Swap，Push，Pop

参考 [Golang 自定义最小堆](https://leetcode-cn.com/problems/top-k-frequent-words/solution/golang-zi-ding-yi-zui-xiao-dui-by-yanert-3zaz/)

### 重要说明：为什么Pop()方法返回数组最后一个元素？

**首先澄清一个误解：数组最后一个元素并不是堆顶！堆顶始终是数组的第一个元素（索引0）。**

我们的Pop()方法返回最后一个元素的原因在于Go标准库heap.Pop()的实现机制：

#### heap.Pop()的工作流程

1. **交换**: 将堆顶元素（索引0，优先级最高）与最后一个元素交换
2. **调整**: 调用down()函数重新调整堆结构（不包括最后一个元素）
3. **移除**: 调用我们实现的Pop()方法移除最后一个元素（此时是原堆顶）并返回

#### 为什么这样设计？

- **避免数组元素大量移动**：如果直接删除堆顶，需要将所有元素前移
- **保持O(log n)时间复杂度**：只需要调整堆结构，不需要移动所有元素
- **高效的删除策略**：通过交换+调整而不是直接删除

#### 堆的基本结构

```text
堆是完全二叉树，用数组表示：
- 堆顶：索引0
- 对于索引i的元素：
  - 左子节点：2*i + 1
  - 右子节点：2*i + 2  
  - 父节点：(i-1) / 2
```

#### 代码中的体现

```go
func (h *WordHeap) Pop() any {
    old := *h
    n := len(old)
    v := old[n-1]  // 返回最后一个元素（已经是原堆顶了）
    *h = old[0 : n-1]
    return v
}
```

所以我们的Pop()方法返回最后一个元素是正确的，因为经过heap.Pop()的处理，最后一个元素已经是原来的堆顶了！

### heap

heap 的 interface

```go
package heap
type Interface interface {
    sort.Interface
    Push(x any) // add x as element Len()
    Pop() any   // remove and return element Len() - 1.
}
```

**注意**: 从 Go 1.18 开始，`any` 是 `interface{}` 的类型别名，推荐使用 `any` 来提高代码可读性。

```go
// Go 1.18+ 中定义的类型别名
type any = interface{}
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




## Solution 2

自己实现的 最小堆, 只包含k个节点。pop出来的k个节点是答案的逆序，需要注意
solution 1 是个最大堆，n个节点（n为数组的size），按照顺序只 pop前k个即可