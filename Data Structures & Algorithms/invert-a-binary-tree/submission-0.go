/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {

	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node == nil {
			return 
		}
		recursion(node.Left)
		recursion(node.Right)
		left := node.Left
		right := node.Right

		node.Left = right 
		node.Right = left
	}

	recursion(root)
	return root
}
