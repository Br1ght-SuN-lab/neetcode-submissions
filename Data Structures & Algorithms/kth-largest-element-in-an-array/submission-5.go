type MinHeap struct {
	data []int
}

func Constructor() MinHeap {
	return MinHeap{
		data: []int{},
	}
}


func (this *MinHeap) Push(value int) {
	this.data = append(this.data, value)
	i := len(this.data)-1
	for i > 0 {
		parent := (i - 1) / 2

		if this.data[parent] < this.data[i] {
			break
		}
		this.data[parent], this.data[i] = this.data[i], this.data[parent]
		i = parent
	}
}


func (this *MinHeap) Pop() {
	this.data[0] = this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]

	i := 0
	smallest := i
	for {
		left, right := 2*i+1, 2*i+2
		if left < len(this.data) && this.data[left] < this.data[smallest] {
			smallest = left
		}
		if right < len(this.data) && this.data[right] < this.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		this.data[smallest], this.data[i] = this.data[i], this.data[smallest]
		i = smallest
	}
}


func (this *MinHeap) Top() int {
	return this.data[0]
}


func findKthLargest(nums []int, k int) int {
	obj := Constructor()
	for _, el := range nums {
		obj.Push(el)
		if len(obj.data) > k {
			obj.Pop() 
		}
	}

	return obj.Top()
}