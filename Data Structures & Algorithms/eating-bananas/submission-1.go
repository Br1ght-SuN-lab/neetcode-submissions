import (
	"slices"
)

func cntOfHours(piles []int, speed int) int {
	hours := 0
	for _, elem := range piles {
		curHours := float64(elem) / float64(speed)
		hours += int(curHours)
		if curHours - float64(int(curHours)) > 1e-9 {
			hours++
		}
	}

	return hours 
}


func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := slices.Max(piles)+1
	ans := 0
	for left < right {
		mid := (left + right) / 2
		result := cntOfHours(piles, mid)
		if result < h {
			right = mid	
			ans = mid	
		} else if result > h {
			left = mid + 1
		} else {
			right = mid
			ans = mid
		}
	}

	return ans
}
