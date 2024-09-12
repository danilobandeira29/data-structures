package main

import "fmt"

/*
  [25, 12, 48, 37, 57, 92, 33]
      0   1   2   3   4   5   6
    down = 1
    up = 4
    fim da primeira iteração do while externo

    array[down] <= pivot and down < up
        down++
    down = 2 = 48 > pivot, entao para

    array[up] > pivot
        up--
    up = 1

    down < up? não, então não troca

    [12, 25, 48, 37, 57, 92, 33]
      0   1   2  3   4    5   6
    segunda iteração do while externo
    quebrou o loop, pois down é maior que up

    trocar start pelo up
    array[start], array[up] = array[up], pivot
*/

var arr = []int{25, 57, 48, 37, 12, 92, 33}

func partition(start, end int) int {
	pivot := arr[start]
	down := start
	up := end
	for down < up {
		for arr[down] <= pivot && down < up {
			down++
		}
		for arr[up] > pivot {
			up--
		}
		if down < up {
			arr[down], arr[up] = arr[up], arr[down]
		}
	}
	arr[start], arr[up] = arr[up], pivot
	return up
}

func QuickSort(start, end int) {
	if start >= end {
		return
	}
	pivot := partition(start, end)
	QuickSort(start, pivot-1) // aplicando para elementos menores que o pivot, antes dele(que não estão ordenados)
	QuickSort(pivot+1, end)   // aplicando para elementos maiores que o pivot, depois dele(que não estão ordenados)
}

func main() {
	QuickSort(0, len(arr)-1)
	fmt.Println(arr)
}
