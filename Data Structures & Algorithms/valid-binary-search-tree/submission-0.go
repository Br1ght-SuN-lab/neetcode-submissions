/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

	var dfs func(node *TreeNode) (bool, int, int)
	dfs = func(node *TreeNode) (bool, int, int) {
		if node == nil {
			return true, math.MaxInt, math.MinInt
		}

		leftValid, leftMin, leftMax := dfs(node.Left)
		rightValid, rightMin, rightMax := dfs(node.Right)

		if !leftValid || !rightValid {
			return false, 0, 0
		}

		if node.Left != nil && leftMax >= node.Val {
			return false, 0, 0
		}

		if node.Right != nil && rightMin <= node.Val {
			return false, 0, 0
		} 

		mn, mx := node.Val, node.Val
		if node.Left != nil {
			mn = leftMin 
		}
		if node.Right != nil {
			mx = rightMax
		}
	
		return true, mn, mx
	}	

	flag, _, _ := dfs(root)
	return flag
}
