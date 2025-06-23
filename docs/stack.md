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

## 单调栈
[单调栈，你该了解的，这里都讲了！LeetCode:739.每日温度](https://www.bilibili.com/video/BV1my4y1Z7jj?vd_source=3d0b9f23ea3cd379598f333513f72565)