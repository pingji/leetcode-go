package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		size := len(stack)
		node := stack[size-1]
		res = append(res, node.Val)
		stack = stack[:size-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// reverse
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return res
}
