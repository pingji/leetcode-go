package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	cur := root

	for cur != nil {
		if abs(float64(cur.Val)-target) < abs(float64(closest)-target) {
			closest = cur.Val
		}
		if target == float64(cur.Val) {
			break
		} else if target < float64(cur.Val) {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return closest
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
