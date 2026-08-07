/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		layerSize := len(queue)
		currentLayer := []int{}

		for i:=0; i < layerSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLayer = append(currentLayer, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLayer)
	}

	return result 
}
