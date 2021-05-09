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
    Node1[2] --a--> Node21[3]
    Node1 --b--> Node22[3]
    Node1 --c--> Node23[3]
    Node21 --d--> End1([ad])
    Node21 --e--> End2([ae])
    Node21 --f--> End3([af])
    Node22 --d--> End4([bd])
    Node22 --e--> End5([be])
    Node22 --f--> End6([bf])
    Node23 --d--> End7([cd])
    Node23 --e--> End8([ce])
    Node23 --f--> End9([cf])
```