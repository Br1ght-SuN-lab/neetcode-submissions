import (
	"slices"
)

type KthLargest struct {
	data []int
	k int
}

func binSearch(data []int, target int) int {
	left := 0
	right := len(data)
	for left < right {
		mid := (left + right) / 2
		if data[mid] < target {
			left = mid + 1
		} else if data[mid] > target {
			right = mid
		} else {
			return mid
		}
	}

	return left
}


func Constructor(k int, nums []int) KthLargest {
	mas := nums
	if len(nums) >= k {
		mas = nums[len(nums)-k:]
	}

	slices.Sort(nums)
	return KthLargest{
		data: mas,
		k: k,
	}
}

func insertVal(data []int, value int) []int {
	result := []int{}
	insertPos := binSearch(data, value)
	left := data[:insertPos]
	right := data[insertPos:]
	result = append(result, left...)
	result = append(result, value)
	result = append(result, right...)

	return result
}


func (this *KthLargest) Add(val int) int {
	if len(this.data) < this.k {
		result := insertVal(this.data, val)
		this.data = result
		return this.data[0]
	}

	if val < this.data[0] {
		return this.data[0]
	}

	result := insertVal(this.data, val)

	this.data = result[1:]
	return this.data[0]
}
