package bubblesort

func Sort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}

func SortFunc[T any](v []T, f func(a, b T) int) {
	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-i-1; j++ {
			result := f(v[j], v[j+1])
			if result > 0 {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}
}
