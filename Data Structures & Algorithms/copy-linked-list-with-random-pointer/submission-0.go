/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	mapOldToNew := make(map[*Node]*Node)
	for p:=head; p!=nil; p = p.Next {
		mapOldToNew[p] = &Node{Val: p.Val} //потом уже будем цеплять 
	}

	//теперь все копии будут подвязаны между собой 
	for node := head; node != nil; node = node.Next {
		copyNode := mapOldToNew[node]
		copyNode.Next = mapOldToNew[node.Next]
		copyNode.Random = mapOldToNew[node.Random]
	}

	return mapOldToNew[head]
}
