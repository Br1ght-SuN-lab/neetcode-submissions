/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	flag := true 
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)

		if math.Abs(float64(left - right)) > 1.0 {
			flag = false 
		}
		
		return max(left, right) + 1 
	}

	depth(root)
	return flag
}
