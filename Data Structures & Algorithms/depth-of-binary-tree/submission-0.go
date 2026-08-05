/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var answer int 
	var recursion func(node *TreeNode, depth int) 
	recursion = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		recursion(node.Left, depth+1)
		answer = max(answer, depth)
		recursion(node.Right, depth+1)
		answer = max(answer, depth)
	} 

	recursion(root, 1)
	return answer
}
