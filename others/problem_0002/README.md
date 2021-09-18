# Solutions

## Solution 1：
找出最长的只包含一个字符的子串

source = “aaabbbcccccdcccdeccc”

Output: “ccccc”


## Solution 2：
找出最长的只包含两个个字符的子串

source = “aaabbbcccccdcccdeccc”

Output: “cccccdcccd”


Solution：增加一个先进先出的队列，保存最近两个不同的字符
# References
