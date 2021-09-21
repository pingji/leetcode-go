package question1

import (
	"fmt"
	"time"
	// ====== 可以引用Golang原生语言包 ====== //
)

// IWorkload 请勿修改接口
type IWorkload interface {
	// Process内包含一些耗时的处理，可能是密集计算或者外部IO
	Process()
}

var TimeoutError = fmt.Errorf("timeout")

// 问题1：请编写函数Question1的实现如下功能
// 该函数输入一个IWorkload实例，请调用其Process函数一次，
// 调用完毕则Question1返回，此时返回的error应为空
// 当Process函数执行5秒仍未能结束时，让Question1函数不再等待
// 立即返回TimeoutError
//
// 注意：题目要求只调用Process一次
// 注意：超时时间固定5秒，请不要修改Question1函数的输入、输出定义
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

func Question1(workload IWorkload) (err error) {
	done := make(chan bool, 1)
	timeout := time.NewTimer(5 * time.Second)
	defer timeout.Stop()

	go func(done chan bool) {
		workload.Process()
		done <- true
		close(done)
	}(done)

	select {
	case <-timeout.C:
		fmt.Print("Exit with timeout")
		return TimeoutError
	case <-done:
		return nil
	}
}
