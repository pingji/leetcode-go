# Solutions

参考 
- [无重复字符的最长子串 c++实现三种解法 多重循环，hashmap优化，桶优化](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/41673/wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-cshi-xian-/)
- [无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/227999/wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-by-leetc-2/)

## Solution 1：滑动窗口 + hashmap

假设我们选择字符串中的第 l 个字符作为起始位置，并且得到了不包含重复字符的最长子串的结束位置为 r。那么当我们选择第 l+1 个字符作为起始位置时，首先从 l+1 到 r的字符显然是不重复的，并且由于少了原本的第 l 个字符，我们可以尝试继续增大r，直到右侧出现了重复字符为止。

这样一来，我们就可以使用「滑动窗口」来解决这个问题了：
