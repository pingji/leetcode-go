package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	curr := root
	for len(stack) > 0 || curr != nil {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			size := len(stack)
			curr = stack[size-1]
			stack = stack[:size-1]
			res = append(res, curr.Val)
			curr = curr.Right
		}
	}

	return res
}
