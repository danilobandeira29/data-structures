package selectionsort

func Sort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(v); j++ {
			if v[j] < v[minIdx] {
				minIdx = j
			}
		}
		if v[minIdx] < v[i] {
			v[minIdx], v[i] = v[i], v[minIdx]
		}
	}
}
