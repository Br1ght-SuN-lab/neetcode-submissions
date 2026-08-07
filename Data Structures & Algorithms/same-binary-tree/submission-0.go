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
func getTreeLayers(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		layerSize := len(queue)

		currentLayer := []int{}
		for i:=0; i < layerSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				currentLayer = append(currentLayer, -1000) //nil-marker
				continue 
			}

			currentLayer = append(currentLayer, node.Val)
			queue = append(queue, node.Left) 
			queue = append(queue, node.Right)
		}
		result = append(result, currentLayer)
	}

	fmt.Println(result)
	return result 
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	firstLayers := getTreeLayers(p)
	secondLayers := getTreeLayers(q)
	if slices.EqualFunc(firstLayers, secondLayers, func(a, b []int) bool {
		return slices.Equal(a, b)
	}) {
		return true
	}
	return false
}
