func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow != fast || (slow == 0 && fast == 0) {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break 
		}
	}

	fmt.Println(slow)

	newSlow := 0
	for newSlow != slow {
		newSlow = nums[newSlow]
		slow = nums[slow]
		if newSlow == slow {
			break 
		}
	}

	return slow
}
