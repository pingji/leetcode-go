package solution

import (
	"fmt"
	"testing"
)

// 定义接口 InterfaceA
type InterfaceA interface {
	Method1()
	Method2()
}

// 测试类型 1：所有方法都是值接收者
type Type1 struct {
	name string
}

func (t Type1) Method1() {
	fmt.Printf("Type1.Method1 called on %s\n", t.name)
}

func (t Type1) Method2() {
	fmt.Printf("Type1.Method2 called on %s\n", t.name)
}

// 测试类型 2：有指针接收者方法
type Type2 struct {
	name string
}

func (t Type2) Method1() {
	fmt.Printf("Type2.Method1 called on %s\n", t.name)
}

func (t *Type2) Method2() { // 指针接收者
	fmt.Printf("Type2.Method2 called on %s\n", t.name)
}

// 测试类型 3：所有方法都是指针接收者
type Type3 struct {
	name string
}

func (t *Type3) Method1() { // 指针接收者
	fmt.Printf("Type3.Method1 called on %s\n", t.name)
}

func (t *Type3) Method2() { // 指针接收者
	fmt.Printf("Type3.Method2 called on %s\n", t.name)
}

// 接受接口 InterfaceA 的函数
func func1(a InterfaceA) {
	fmt.Println("=== func1 调用 ===")
	a.Method1()
	a.Method2()
}

// 演示函数
func DemoInterfaceCall() {
	fmt.Println("=== Go 接口调用规则演示 ===")

	// 测试 Type1 (所有方法都是值接收者)
	fmt.Println("1. Type1 (所有方法都是值接收者):")
	t1 := Type1{name: "type1"}

	fmt.Println("   func1(t1):")
	func1(t1) // ✓ 可以

	fmt.Println("   func1(&t1):")
	func1(&t1) // ✓ 也可以

	fmt.Println()

	// 测试 Type2 (有指针接收者方法)
	fmt.Println("2. Type2 (有指针接收者方法):")
	t2 := Type2{name: "type2"}

	fmt.Println("   func1(t2):")
	// func1(t2)  // ✗ 编译失败！Type2 没有实现接口 InterfaceA
	fmt.Println("   // func1(t2) 会编译失败！")

	fmt.Println("   func1(&t2):")
	func1(&t2) // ✓ 可以

	fmt.Println()

	// 测试 Type3 (所有方法都是指针接收者)
	fmt.Println("3. Type3 (所有方法都是指针接收者):")
	t3 := Type3{name: "type3"}

	fmt.Println("   func1(t3):")
	// func1(t3)  // ✗ 编译失败！Type3 没有实现接口 InterfaceA
	fmt.Println("   // func1(t3) 会编译失败！")

	fmt.Println("   func1(&t3):")
	func1(&t3) // ✓ 可以

	fmt.Println()

	// 接口实现检查
	fmt.Println("=== 接口实现检查 ===")

	// Type1 的检查
	var _ InterfaceA = t1  // ✓ Type1 实现了接口 InterfaceA
	var _ InterfaceA = &t1 // ✓ *Type1 也实现了接口 InterfaceA
	fmt.Println("✓ Type1 实现了接口 InterfaceA")
	fmt.Println("✓ *Type1 也实现了接口 InterfaceA")

	// Type2 的检查
	// var _ InterfaceA = t2   // ✗ 编译失败！Type2 没有实现接口 InterfaceA
	var _ InterfaceA = &t2 // ✓ *Type2 实现了接口 InterfaceA
	fmt.Println("✗ Type2 没有实现接口 InterfaceA")
	fmt.Println("✓ *Type2 实现了接口 InterfaceA")

	// Type3 的检查
	// var _ InterfaceA = t3   // ✗ 编译失败！Type3 没有实现接口 InterfaceA
	var _ InterfaceA = &t3 // ✓ *Type3 实现了接口 InterfaceA
	fmt.Println("✗ Type3 没有实现接口 InterfaceA")
	fmt.Println("✓ *Type3 实现了接口 InterfaceA")
}

// 接口演示测试
func TestInterfaceDemo(t *testing.T) {
	DemoInterfaceCall()
}

// 方法集详细分析
func TestMethodSet(t *testing.T) {
	fmt.Println("=== 方法集分析 ===")

	// Type1 的方法集
	fmt.Println("Type1 的方法集:")
	fmt.Println("- func (t Type1) Method1()")
	fmt.Println("- func (t Type1) Method2()")
	fmt.Println("结论: Type1 有 Method1 和 Method2 → 实现了 InterfaceA")

	fmt.Println("\n*Type1 的方法集:")
	fmt.Println("- func (t Type1) Method1()  (从 Type1 提升)")
	fmt.Println("- func (t Type1) Method2()  (从 Type1 提升)")
	fmt.Println("结论: *Type1 也有 Method1 和 Method2 → 实现了 InterfaceA")

	fmt.Println("\n" + repeatString("=", 50))

	// Type2 的方法集
	fmt.Println("Type2 的方法集:")
	fmt.Println("- func (t Type2) Method1()   ✓ 有")
	fmt.Println("- func (t *Type2) Method2()  ✗ 没有 (这是指针接收者)")
	fmt.Println("结论: Type2 缺少 Method2 → 没有实现 InterfaceA")

	fmt.Println("\n*Type2 的方法集:")
	fmt.Println("- func (t Type2) Method1()   ✓ 有 (从 Type2 提升)")
	fmt.Println("- func (t *Type2) Method2()  ✓ 有 (指针接收者)")
	fmt.Println("结论: *Type2 有 Method1 和 Method2 → 实现了 InterfaceA")

	fmt.Println("\n" + repeatString("=", 50))

	// Type3 的方法集
	fmt.Println("Type3 的方法集:")
	fmt.Println("- func (t *Type3) Method1()  ✗ 没有 (这是指针接收者)")
	fmt.Println("- func (t *Type3) Method2()  ✗ 没有 (这是指针接收者)")
	fmt.Println("结论: Type3 什么方法都没有 → 没有实现 InterfaceA")

	fmt.Println("\n*Type3 的方法集:")
	fmt.Println("- func (t *Type3) Method1()  ✓ 有 (指针接收者)")
	fmt.Println("- func (t *Type3) Method2()  ✓ 有 (指针接收者)")
	fmt.Println("结论: *Type3 有 Method1 和 Method2 → 实现了 InterfaceA")
}

func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
