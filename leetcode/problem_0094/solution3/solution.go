package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node != nil {
			// pop
			// 将该节点弹出，避免重复操作，下面再将右中左节点添加到栈中
			stack = stack[:len(stack)-1]

			// 1. 右
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			// 2. 中
			stack = append(stack, node)
			// 中节点访问过，但是还没有处理，加入空节点做为标记
			stack = append(stack, nil)
			// 3. 左
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			// 只有遇到空节点的时候，才将下一个节点放进结果集
			// pop nil
			stack = stack[:len(stack)-1]

			node = stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]
			// output
			res = append(res, node.Val)
		}
	}

	return res
}
