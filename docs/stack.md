## 如何实现 stack
golang中实现stack的功能

1. 使用slice定义stack
```
stack := []int{}
```

2. push
```
stack = append（stack, value)
```
3. top
```
value := stack[len(stack)-1]
```
4. pop
```
stack = stack[:len(stack)-1]
```