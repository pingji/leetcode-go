package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		currLevelSize := len(q)
		currLevelRes := []int{}
		for i := 0; i < currLevelSize; i++ {
			node := q[0]
			q = q[1:]
			currLevelRes = append(currLevelRes, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append([][]int{currLevelRes}, res...)
	}
	return res
}
