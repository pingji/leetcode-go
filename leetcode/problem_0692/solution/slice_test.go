package solution

import (
	"fmt"
	"testing"
)

type TestSlice []int

// 值接收者 - 不会修改原切片的长度
func (s TestSlice) AppendValue(x int) {
	// 本质：s 是原切片的一个副本
	// 切片结构：{ptr: 指向底层数组的指针, len: 长度, cap: 容量}
	// 这里的 s 是新的切片结构，但 ptr 指向同一个底层数组
	s = append(s, x)
	fmt.Printf("值接收者内部: %v\n", s)
}

// 指针接收者 - 会修改原切片的长度
func (s *TestSlice) AppendPointer(x int) {
	// 本质：s 是指向原切片的指针
	// 通过 *s 直接修改原切片的结构
	*s = append(*s, x)
	fmt.Printf("指针接收者内部: %v\n", *s)
}

// 值接收者 - 可以修改切片中的元素
func (s TestSlice) SwapElements(i, j int) {
	// 关键理解：虽然 s 是切片的副本，但是：
	// 1. 切片结构被复制了（新的 len, cap, ptr）
	// 2. 但 ptr 指向的是同一个底层数组！
	// 3. 通过 s[i] 访问的是底层数组的元素，不是切片结构本身
	// 4. 所以可以修改底层数组的内容，但不能修改切片的结构（len, cap）

	fmt.Printf("交换前: s[%d]=%d, s[%d]=%d\n", i, s[i], j, s[j])
	s[i], s[j] = s[j], s[i] // 修改底层数组的元素
	fmt.Printf("交换后内部: %v\n", s)
}

func TestSliceReceiver(t *testing.T) {
	slice := TestSlice{1, 2, 3}
	fmt.Printf("原始切片: %v\n", slice)

	fmt.Println("\n--- 测试值接收者添加元素 ---")
	slice.AppendValue(4)
	fmt.Printf("调用后外部: %v\n", slice) // 长度没变

	fmt.Println("\n--- 测试指针接收者添加元素 ---")
	slice.AppendPointer(4)
	fmt.Printf("调用后外部: %v\n", slice) // 长度改变了

	fmt.Println("\n--- 测试值接收者交换元素 ---")
	slice.SwapElements(0, 1)
	fmt.Printf("调用后外部: %v\n", slice) // 元素被交换了
}

// 添加一个测试来显示内存地址差异
func TestSliceMemoryAddress(t *testing.T) {
	slice := TestSlice{1, 2, 3}
	fmt.Printf("原始切片: %v\n", slice)
	fmt.Printf("原始切片地址: %p\n", &slice)
	fmt.Printf("原始切片底层数组地址: %p\n", &slice[0])

	fmt.Println("\n--- 值接收者内存分析 ---")
	slice.ShowValueReceiverMemory()

	fmt.Println("\n--- 指针接收者内存分析 ---")
	slice.ShowPointerReceiverMemory()

	fmt.Printf("调用后原切片: %v\n", slice)
	fmt.Printf("调用后原切片地址: %p\n", &slice)
}

// 值接收者 - 查看内存地址
func (s TestSlice) ShowValueReceiverMemory() {
	fmt.Printf("值接收者内部切片地址: %p\n", &s)
	fmt.Printf("值接收者内部底层数组地址: %p\n", &s[0])
	fmt.Printf("值接收者内部切片: %v\n", s)
}

// 指针接收者 - 查看内存地址
func (s *TestSlice) ShowPointerReceiverMemory() {
	fmt.Printf("指针接收者s地址: %p\n", s)
	fmt.Printf("指针接收者*s地址: %p\n", s) // s 就是指向切片的指针
	fmt.Printf("指针接收者底层数组地址: %p\n", &(*s)[0])
	fmt.Printf("指针接收者内部切片: %v\n", *s)
}

// 详细展示切片结构的测试
func TestSliceStructure(t *testing.T) {
	slice := TestSlice{1, 2, 3}

	fmt.Printf("=== 切片结构分析 ===\n")
	fmt.Printf("原始切片: %v\n", slice)
	fmt.Printf("原始切片地址: %p\n", &slice)
	fmt.Printf("原始切片底层数组地址: %p\n", &slice[0])
	fmt.Printf("原始切片长度: %d, 容量: %d\n", len(slice), cap(slice))

	fmt.Printf("\n=== 值接收者修改元素 ===\n")
	slice.DetailedSwapElements(0, 1)

	fmt.Printf("\n=== 修改后的原切片 ===\n")
	fmt.Printf("修改后切片: %v\n", slice)
	fmt.Printf("修改后切片地址: %p\n", &slice)
	fmt.Printf("修改后切片底层数组地址: %p\n", &slice[0])
	fmt.Printf("修改后切片长度: %d, 容量: %d\n", len(slice), cap(slice))
}

// 详细展示交换过程的方法
func (s TestSlice) DetailedSwapElements(i, j int) {
	fmt.Printf("进入值接收者方法：\n")
	fmt.Printf("  参数切片地址: %p\n", &s)
	fmt.Printf("  参数切片底层数组地址: %p\n", &s[0])
	fmt.Printf("  参数切片长度: %d, 容量: %d\n", len(s), cap(s))
	fmt.Printf("  参数切片内容: %v\n", s)

	fmt.Printf("执行交换 s[%d] ↔ s[%d]：\n", i, j)
	fmt.Printf("  交换前: s[%d]=%d, s[%d]=%d\n", i, s[i], j, s[j])

	s[i], s[j] = s[j], s[i] // 这里修改的是底层数组

	fmt.Printf("  交换后: s[%d]=%d, s[%d]=%d\n", i, s[i], j, s[j])
	fmt.Printf("  交换后切片内容: %v\n", s)
}

/*
=== 切片的内部结构解析 ===

Go 切片的内部结构：
┌─────────────────────────────────────────────────────────────────┐
│                        切片 (slice)                             │
├─────────────────────────────────────────────────────────────────┤
│  ptr: 0x140001ca288  (指向底层数组的指针)                        │
│  len: 3              (当前长度)                                  │
│  cap: 3              (容量)                                      │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ 指向
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      底层数组 (array)                           │
├─────────────────────────────────────────────────────────────────┤
│  [0]: 1    [1]: 2    [2]: 3                                     │
│  0x140001ca288 + 0*8  0x140001ca288 + 1*8  0x140001ca288 + 2*8  │
└─────────────────────────────────────────────────────────────────┘

=== 值接收者调用过程 ===

原始切片 (地址: 0x140001fc180):
┌─────────────────────────────────────────────────────────────────┐
│  ptr: 0x140001ca288  │  len: 3  │  cap: 3                       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ 复制切片结构
                              ▼
参数切片 (地址: 0x140001fc1b0):
┌─────────────────────────────────────────────────────────────────┐
│  ptr: 0x140001ca288  │  len: 3  │  cap: 3                       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ 两个切片的 ptr 指向同一个底层数组！
                              ▼
                    底层数组 (地址: 0x140001ca288):
                    ┌─────────────────────────────────────┐
                    │  [0]: 1→2  [1]: 2→1  [2]: 3        │
                    └─────────────────────────────────────┘
                              ▲
                              │ s[i], s[j] = s[j], s[i]
                              │ 修改的是底层数组的内容

=== 为什么值接收者可以修改元素？ ===

1. **切片结构被复制**：新的切片有自己的 len, cap, ptr 变量
2. **但 ptr 指向同一个底层数组**：两个切片共享同一个底层数组
3. **s[i] 访问的是底层数组**：通过索引访问的是底层数组的元素，不是切片结构本身
4. **因此可以修改底层数组的内容**：元素的修改会影响所有指向该底层数组的切片

=== 为什么值接收者不能修改切片长度？ ===

1. **append 操作修改的是切片结构**：len 和 cap 字段，以及可能的 ptr 重新分配
2. **值接收者的切片结构是副本**：对副本的修改不会影响原切片
3. **因此 append 的结果只在方法内部可见**

这就是为什么：
- Swap() 可以用值接收者：只修改底层数组元素
- Push()/Pop() 必须用指针接收者：需要修改切片结构
*/
