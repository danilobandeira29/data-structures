package insertionsort

// this works too
// func Sort(v []int) {
// 	for i := 1; i < len(v); i++ {
// 		for temp := i; temp >= 1 && v[temp] < v[temp-1]; temp-- {
// 			v[temp], v[temp-1] = v[temp-1], v[temp]
// 		}
// 	}
// }

func SortFunc[T any](v []T, f func(a, b T) int) {
	for i := 1; i < len(v); i++ {
		temp := i
		for temp >= 1 && f(v[temp-1], v[temp]) > 0 {
			v[temp], v[temp-1] = v[temp-1], v[temp]
			temp--
		}
	}
}

func Sort(v []int) {
	for i := 1; i < len(v)-1; i++ {
		card := v[i]
		j := i - 1
		for ; j >= 0 && v[j] > card; j-- {
			v[j+1] = v[j]
		}
		v[j+1] = card
	}
}
