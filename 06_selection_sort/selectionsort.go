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

func SortFunc[T any](v []T, f func(a, b T) int) {
	for i := 0; i < len(v)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(v); j++ {
			result := f(v[minIdx], v[j])
			if result > 0 {
				minIdx = j
			}
		}
		result := f(v[i], v[minIdx])
		if result > 0 {
			v[minIdx], v[i] = v[i], v[minIdx]
		}
	}
}
