func binSearch(data []float64, value float64) int {
	left := 0
	right := len(data)
	eps := 1e-8
	for left < right {
		mid := (left + right) / 2
		if data[mid] < value {
			left = mid + 1
		} else if data[mid] > value {
			right = mid
		} else if data[mid]-value < eps {
			return mid
		}
	}
	return left
}

func insertValue(data []float64, value float64) []float64 {
	if len(data) == 0 {
		return []float64{value}
	}
	insertPos := binSearch(data, value)
	result := []float64{}
	left := data[:insertPos]
	right := data[insertPos:]
	result = append(result, left...)
	result = append(result, value)
	result = append(result, right...)
	return result
}

func fillFields(distances *[]float64, mapa *map[float64][][]int, dist float64, point []int) {
	*distances = insertValue(*distances, dist)
	(*mapa)[dist] = append((*mapa)[dist], point)
}

func kClosest(points [][]int, k int) [][]int {
	distances := []float64{}
	mapa := make(map[float64][][]int, 0)
	for _, point := range points {
		dist := math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
		if len(distances) >= k && dist >= distances[k-1] {
			continue
		}

		fillFields(&distances, &mapa, dist, point)

		if len(distances) > k {
			worst := distances[k]
			distances = distances[:k]
			pts := mapa[worst]
			if len(pts) > 1 {
				mapa[worst] = pts[:len(pts)-1]
			} else {
				delete(mapa, worst)
			}
		}
	}

	answer := [][]int{}
	for _, elems := range mapa {
		for _, point := range elems {
			answer = append(answer, point)
		}
	}

	return answer
}