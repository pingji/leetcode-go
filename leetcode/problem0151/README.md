## Solutions
### solution 1
1. 利用 strings的 Filed 去掉多余空格并分割单词
2. 倒序放入另一个字符串数组
3. 将倒序的字符串数组 Join起来
   
### solution 2
因为golang的string无法改变，需要使用 []byte(s) 将string转成[]，这个过程发生了拷贝
1. 先去掉多余的空格
2. 字符串反转
3. 单词反转

### solution 3
使用空间复杂度 o(1) 的方案，因为golang的string无法改变，需要使用unsafe.Pointer将string转成[]byte
1. 先去掉多余的空格
2. 字符串反转
3. 单词反转
使用golang-1.16.2编译，运行时发生“unexpected fault address” 异常，但是在leetcode上提交正常运行