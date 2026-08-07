import (
	"slices"
)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		layerSize := len(queue)
		currentLayer := []int{}
		for i := 0; i < layerSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				currentLayer = append(currentLayer, node.Val)
				queue = append(queue, node.Right)
				queue = append(queue, node.Left)
			}
		}
		if len(currentLayer) == 0 {
			continue
		}
		slices.Reverse(currentLayer)
		result = append(result, currentLayer[len(currentLayer)-1])
	}

	return result 
}
