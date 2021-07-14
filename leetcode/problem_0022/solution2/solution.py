#!/usr/bin/python3
from typing import List

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        if n <= 0: return []
        res = []

        def dfs(paths, left, right):
            if left > n or right > left: return
            if len(paths) == n * 2:  # 因为括号都是成对出现的
                res.append(paths)
                return

            dfs(paths + '(', left + 1, right)
            dfs(paths + ')', left, right + 1)

        dfs('',0 ,0)
        return res

if __name__ == '__main__':
    solution = Solution()
    res = solution.generateParenthesis(3)
    print(res)

