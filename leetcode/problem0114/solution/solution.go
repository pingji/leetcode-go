package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 1、左右子树已经被拉平成一条链表
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right

	// 2、将左子树作为右子树
	root.Right = left
	root.Left = nil

	// 3、将原先的右子树接到当前右子树的末端
	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = right
}
