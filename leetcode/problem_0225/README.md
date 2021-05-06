# Solutions

## Solution 1：双栈

用两个队列实现栈. 一个队列为主队列，一个为辅助队列，当入栈操作时，我们先将主队列内容导入辅助队列，然后将入栈元素放入主队列队头位置，再将辅助队列内容，依次添加进主队列即可。

参考 [用队列实现栈](https://leetcode-cn.com/problems/implement-stack-using-queues/solution/wu-tu-guan-fang-tui-jian-ti-jie-yong-dui-63d4/)