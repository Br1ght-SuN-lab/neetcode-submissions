type MaxHeap []int
func (h MaxHeap) Len() int {return len(h)}

func (h MaxHeap) Less(i, j int) bool {return h[i] > h[j]}

func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *MaxHeap) Push(val any) {*h = append(*h, val.(int))}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int, 0)
	for _, el := range tasks {
		freq[el]++
	}

	h := &MaxHeap{}
	for _, count := range freq {
		heap.Push(h, count)
	}

	time := 0
	queue := [][2]int{}
	for len(queue) != 0 || len(*h) != 0 {
		time++
		if len(*h) != 0 {
			count := heap.Pop(h).(int)-1
			if count != 0 {
				queue = append(queue, [2]int{count, time+n})
			}
		}

		if len(queue) != 0 && queue[0][1] == time {
			heap.Push(h, queue[0][0])
			queue = queue[1:]
		}
	}

	return time
}
