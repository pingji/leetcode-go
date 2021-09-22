# Solutions
```go
package question2

import (
	// ====== 可以引用Golang原生语言包 ====== //
)

// IWorkload 请勿修改接口
type IWorkload interface {
	// Work内包含一些耗时的处理，可能是密集计算或者外部IO
	Work()
}

// IProducer 请勿修改接口
type IProducer interface {
	// Produce每次调用会返回一个IWorkload实例
	// 当返回nil时表示已经生产完毕
	Produce() IWorkload
}

// 问题2：请编写函数Question2的实现如下功能
// 该函数输入一个IProducer实例，每次调用其Produce()方法会返回一个IWorkload实例。
// 1. 请反复调用该Produce()方法，直到返回nil，表明没有更多IWorkload。
//    此间可能会生产大量IWorkload实例，数目在此未知。
// 2. 对每个生产出的IWorkload实例，请调用一次它的Work()方法。
//    Work()内包含一些耗时的处理，可能是密集计算或者外部IO。
// 3. 请并发调用多个IWorkload的Work()方法，最多允许5个并发的Work()执行。
//    单个并发的实现，或并发数超过5的限制，都不能得分。
//
// 提示：请最小化内存、CPU代价
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

func Question2(producer IProducer) {

	// ====== 在这里书写代码 ====== //

}


```


## Solution 1：


# References
