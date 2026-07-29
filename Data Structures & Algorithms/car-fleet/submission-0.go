import (
	"slices"
)

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]int, n)
	for i := range n {
		cars[i] = [2]int{position[i], speed[i]}
	}	

	slices.SortFunc(cars, func(i, j [2]int) int {
		return j[0] - i[0]
	})

	stack := []float64{}
	for _,el := range cars {
		time := float64(target - el[0]) / float64(el[1])
		if len(stack) == 0 {
			stack = append(stack, time)
			continue
		}
		top := stack[len(stack)-1]
		if top < time {
			stack = append(stack, time)
		}
	}


	return len(stack) 
}
