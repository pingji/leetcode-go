# 1. 概要
二叉树是每个树节点最多有两个子树的一种特殊的树结构，其有一些内在的性质，比如，若二叉树的层次从0开始，则在二叉树的第i层至多有 $2^i$个节点（i>=0）,高度为k的二叉树最多有 $2^{k+1}-1$ 个节点（空树的高度为-1）。其类别为以下几种：

# 2. 分类
- 满二叉树：所有的叶节点全部在底层，并且在底层全部铺满的二叉树
- 完全二叉树：叶节点只能出现在最后两层，并且最底层的叶节点都向左对齐
- 二叉搜索树：要求每个节点本身大于其左子树，而小于其右子树，对其进行中序遍历后，会得到一个有序的列表，这是我们经常用到的一种数的结构
- 平衡二叉树：它是一 棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树，并且满足二叉搜索树的规则
# 3. 遍历
## 3.1. 应用场景
- 前序遍历：第一次访问就输出数据，适合静态访问
- 中序遍历：对于二分搜索树，输出结果是有序的，适合顺序输出结果
- 后续遍历：对节点操作时必访问过其子节点，适合进行破坏性操作（删除节点）
  
## 3.2. 遍历的方法
- 递归
- 迭代
### 3.2.1. 递归

### 3.2.2. 迭代
任何算法的递归版本都可以改成非递归版本，因为函数递归调用其实质就是压栈的过程，那么我们完全可以使用堆栈来模拟这个过程

## 3.3. 参考
- [彻底吃透前中后序递归法和迭代法](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/bang-ni-dui-er-cha-shu-bu-zai-mi-mang-che-di-chi-t/)