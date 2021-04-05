## Solutions
### solution1
采用回溯算法,回溯算法的框架为
```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```
需要注意的是，result.add(路径) 需要深度拷贝


```mermaid
graph TB
    root[ ] --> Node11[1]
    root --> Node12[2]
    root --> Node13[3]
    Node11 --> Node21[2]
    Node11 --> Node22[3]
    Node12 --> Node23[1]
    Node12 --> Node24[3]
    Node13 --> Node25[1]
    Node13 --> Node26[2]
    Node21 --> Node31[3]
    Node22 --> Node32[2]
    Node23 --> Node33[3]
    Node24 --> Node34[1]
    Node25 --> Node35[2]
    Node26 --> Node36[1]
```