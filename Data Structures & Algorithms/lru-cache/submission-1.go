type Record struct {
	key, val int
	next, prev *Record
}

//используем dummy подход с двух сторон, потому что двусвязанный список
type LRUCache struct {
	capacity int
	data map[int]*Record
	head *Record
	tail *Record
}

func addNode(head *Record, node *Record) {
	node.next = head.next
	node.prev = head
	head.next.prev = node
	head.next = node
}


func removeNode(node *Record) {
	node.prev.next = node.next
	node.next.prev = node.prev
}


func Constructor(capacity int) LRUCache {
	head := &Record{}
	tail := &Record{}
	
	head.next = tail
	tail.prev = head
	
	return LRUCache {
		capacity: capacity,
		data: make(map[int]*Record, capacity),
		head: head,
		tail: tail, 
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		removeNode(node)
		addNode(this.head, node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		removeNode(node)
		node.val = value
		this.data[key] = node
		addNode(this.head, node)
		return
	}

	if len(this.data) == this.capacity {
		oldNode := this.tail.prev
		removeNode(oldNode)
		delete(this.data, oldNode.key)
	}

	newNode := &Record{key: key, val: value}
	addNode(this.head, newNode)
	this.data[key] = this.head.next
}