package bubblesort

func Sort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] <= values[j+1] {
				continue
			}
			temp := values[j]
			values[j] = values[j+1]
			values[j+1] = temp
		}
	}
	return values
}
