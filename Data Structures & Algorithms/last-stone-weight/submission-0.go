import (
	"slices"
)

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

func insertValue(data []int, value int) []int {
	result := []int{}
	insertPos := binSearch(data, value)
	left := data[:insertPos]
	right := data[insertPos:]
	result = append(result, left...)
	result = append(result, value)
	result = append(result, right...)
	return result 
}


func lastStoneWeight(stones []int) int {
	slices.Sort(stones)

	for len(stones) >= 2 {
		first := stones[len(stones)-1]
		second := stones[len(stones)-2]
		if first == second {
			stones = stones[:len(stones)-2]
			continue
		}

		stones = insertValue(stones[:len(stones)-2], first-second)
		fmt.Println(stones)
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
