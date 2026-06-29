func hasDuplicate(nums []int) bool {
    container := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if container[nums[i]] > 0 {
            return true
        }
        container[nums[i]]++
    }
    return false  
}
