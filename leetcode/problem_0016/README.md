# Solutions

## Solution 1：排序 + 双指针
1. 先排序
2. 固定第一个值，计算三数之和，更新和的绝对值是否与target更近
3. 移动第二第三个值的指针
   1. 如果 nums[p₁] + nums[p₂] + nums[p₃] ≥ target，那么就将 p₃ 向左移动一个位置；
   2. 如果 nums[p₁] + nums[p₂] + nums[p₃] < target，那么就将 p₂ 向右移动一个位置；
4. 注意去重

# References
- [秒杀所有求和问题](https://leetcode-cn.com/problems/3sum/solution/yi-miao-jiu-neng-kan-dong-de-dong-tu-jie-unfp/)