/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode { 
	//для обращения за O(1) к каждому узлу
	inorderMapa := make(map[int]int, len(inorder)+1)
	for i, v := range inorder {
		inorderMapa[v] = i
	}

	//обход по preorder 
	preInd := 0

	var build func(l, r int) *TreeNode 
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil 
		}

		rootVal := preorder[preInd]
		preInd++
		//получаем индекс очередного корня в inorder
		mid := inorderMapa[rootVal]

		node := &TreeNode{Val: rootVal}
		//подвязываем правое и левое поддерево
		node.Left = build(l, mid-1)
		node.Right = build(mid+1, r)
		return node
	}

	return build(0, len(inorder)-1)
}
