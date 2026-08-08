/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	result := 0
	var dfs func(node *TreeNode, maxSoFar int)
	dfs = func(node *TreeNode, maxSoFar int) {
		if node == nil {
			return 
		}

		if node.Val >= maxSoFar {
			maxSoFar = node.Val
			result++
		}

		dfs(node.Left, maxSoFar)
		dfs(node.Right, maxSoFar)
	}

	dfs(root, int(math.Inf(-1)))
	return result
}
