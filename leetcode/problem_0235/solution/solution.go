package solution

import (
	"leetcode-go/common/binary_tree"
)

type TreeNode = binary_tree.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor := root
	for ancestor != nil {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			break
		}
	}
	return ancestor
}
