func twoSum(nums []int, target int) []int {
    previousValues := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        need := target - nums[i]
        if j, ok := previousValues[need]; ok {
            return []int{j, i} //j встретили раньше, поэтому такой порядок
        }
        previousValues[nums[i]] = i
    }

    return nil 
}
