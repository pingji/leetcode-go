package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Input 定义测试输入结构
type Input struct {
	nums  []int // 输入数组
	lower int   // 区间和的下界
	upper int   // 区间和的上界
}

// TestProblem 测试区间和个数计算函数
func TestProblem(t *testing.T) {
	assert := assert.New(t)

	// 定义测试用例
	tests := []struct {
		input  Input // 输入参数
		output int   // 期望输出
	}{
		{
			// 测试用例1：基本测试
			input: Input{
				nums:  []int{-2, 5, -1}, // 输入数组
				lower: -2,               // 下界
				upper: 2,                // 上界
			},
			output: 3, // 期望结果：3个满足条件的区间
			// 解释：
			// [0,0] -> -2 (满足条件)
			// [2,2] -> -1 (满足条件)
			// [0,2] -> 2  (满足条件)
		},
		{
			// 测试用例2：边界情况
			input: Input{
				nums:  []int{0}, // 只有一个元素的数组
				lower: 0,        // 下界
				upper: 0,        // 上界
			},
			output: 1, // 期望结果：1个满足条件的区间
			// 解释：只有区间 [0,0] 的和为 0，满足条件
		},
	}

	// 执行测试用例
	for index, test := range tests {
		// 调用被测试的函数
		output := countRangeSum(test.input.nums, test.input.lower, test.input.upper)

		// 输出测试信息
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)

		// 断言测试结果
		assert.Equal(test.output, output)
	}
}
