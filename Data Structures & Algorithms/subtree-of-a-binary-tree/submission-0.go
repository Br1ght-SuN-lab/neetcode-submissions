/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func treeEqual(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true 
	} 
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return treeEqual(p.Left, q.Left) && treeEqual(p.Right, q.Right)
} 


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	fmt.Println(subRoot)
    var recursion func(node *TreeNode) bool
	recursion = func(node *TreeNode) bool {
		fmt.Println(node)
		if node == nil {
			return false
		}

		if treeEqual(node, subRoot) {
			return true
		}

		left := recursion(node.Left)
		if left {
			return true
		}
		right := recursion(node.Right)
		if right {
			return true
		}

		return false 
	}

	result := recursion(root)
	return result 
}
