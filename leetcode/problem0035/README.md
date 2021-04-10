## Solutions
### solution 1
二分查找，找到了则返回mid，找不到则返回left

考虑找不到的情况
- 当 target 比 left 还小，right 会跳到 left的左边，返回left
- 当 tatget 逼 right 还大，left 会跳到 right的右边，返回left