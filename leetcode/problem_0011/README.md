# Solutions

## Solution 1：双指针
S(i,j)=min(h[i],h[j])×(j−i)

在每个状态下，无论长板或短板向中间收窄一格，都会导致水槽 底边宽度 −1变短：
- 若向内 移动短板 ，水槽的短板 min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
- 若向内 移动长板 ，水槽的短板 min(h[i],h[j]) 不变或变小，因此下个水槽的面积 一定变小。

因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积。

算法流程：
1. 初始化： 双指针 i , j 分列水槽左右两端；
2. 循环收窄： 直至双指针相遇时跳出；
    1. 更新面积最大值 res ；
    2. 选定两板高度中的短板，向中间收窄一格；
3. 返回值： 返回面积最大值 res 即可；


# References
- [盛最多水的容器（双指针，清晰图解）](https://leetcode.cn/problems/container-with-most-water/solutions/11491/container-with-most-water-shuang-zhi-zhen-fa-yi-do/)