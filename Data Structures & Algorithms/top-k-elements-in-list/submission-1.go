func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, elem := range nums {
		freq[elem]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)

	fmt.Println(buckets)

	flag := true
	for i := len(nums); i >= 0; i-- {
		for _, elem := range buckets[i] {
			if len(result) >= k {
				flag = false
				break
			}

			result = append(result, elem)
		}

		if flag == false {
			break
		}
	}

	return result 
}