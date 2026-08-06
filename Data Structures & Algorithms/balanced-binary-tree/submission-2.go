/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		//поднимаем наверх несбалансированность
		if left == -1 {
			return -1
		}
		right := depth(node.Right)
		if right == -1 {
			return -1
		}

		if math.Abs(float64(left - right)) > 1.0 {
			return -1 
		}
		
		return max(left, right) + 1 
	}

	result := depth(root)
	return result != -1
}
