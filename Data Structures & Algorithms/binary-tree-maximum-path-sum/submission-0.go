/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	result := root.Val

	var dfs func(node *TreeNode) int 
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)

		result = max(result, node.Val + leftMax + rightMax)
		return max(rightMax, leftMax) + node.Val
	}

	dfs(root)
	return result
}
