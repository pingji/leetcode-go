# Solutions
## Solution 1: 
引入队列，与102不同的是，每一行的列表不是append在尾部，而是insert到头部

golang中，slice的append方法
```go
var res [][]int
sum:=[]int{2}
res=append(res,sum)
```
golang中，slice的insert 头部方法
```go
var res [][]int
sum:=[]int{2}
res=append([][]int{sum},res...)
```