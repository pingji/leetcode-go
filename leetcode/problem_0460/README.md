# Solutions

## Solution 1: 双哈希表 + 双向链表
- 第一个哈希表 key - node，
- 第二个哈希表 freq - node list （双向链表）

技巧
1. 选用双向链表目的是为了达到 删除链表中节点的时间复杂度可以做到 O(1)，如果不是O(1)，测试用例会提示“超出时间限制”
2. 在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。
3. 维护一个 min_freq，记录freq map中最小的 freq，当cache满时，删除min_freq中最久未访问的node
