func binarySearch(keys []int, target int) int {
	fmt.Println(keys)
	left := 0
	right := len(keys)-1
	for left <= right {
		mid := (left + right) / 2
		if keys[mid] > target {
			right = mid-1
		} else if keys[mid] < target {
			left = mid+1
		} else {
			return keys[mid]
		}
	}

	//случай если timestratch меньше чем возможно
	if right == -1 {
		return -1 
	}
	return keys[right]
} 


type TimeMap struct {
	data map[string]map[int]string
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string]map[int]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.data[key] == nil {
		this.data[key] = make(map[int]string)
	}
	this.data[key][timestamp] = value
}


func (this *TimeMap) Get(key string, timestamp int) string {
	innerKeys := []int{}
	for k, _ := range this.data[key] {
		innerKeys = append(innerKeys, k)
	}

	sort.Ints(innerKeys)
	
	needInnerKey := binarySearch(innerKeys, timestamp)
	if needInnerKey == -1 {
		return ""
	}
	return this.data[key][needInnerKey]
}
