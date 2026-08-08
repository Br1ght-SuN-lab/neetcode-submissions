/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	result := -1
	var dfs func(node *TreeNode)	
	dfs = func(node *TreeNode) {
		if node == nil {
			return 
		}


		dfs(node.Left)
		k-- 
		if k == 0 && result == -1 {
			result = node.Val
		}
		dfs(node.Right)
	}

	dfs(root)
	return result 
}
