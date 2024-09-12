package quicksort

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

type quickSort struct {
	values []int
}

func (q *quickSort) partition(start, end int) int {
	pivot := q.values[start]
	down := start
	up := end
	for down < up {
		for q.values[down] <= pivot && down < up {
			down++
		}
		for q.values[up] > pivot {
			up--
		}
		if down < up {
			q.values[down], q.values[up] = q.values[up], q.values[down]
		}
	}
	q.values[start], q.values[up] = q.values[up], pivot
	return up
}

func (q *quickSort) do(start, end int) {
	if start >= end {
		return
	}
	pivot := q.partition(start, end)
	q.do(start, pivot-1) // aplicando para elementos menores que o pivot, antes dele(que não estão ordenados)
	q.do(pivot+1, end)   // aplicando para elementos maiores que o pivot, depois dele(que não estão ordenados)
}

func Sort(v []int) {
	q := &quickSort{
		values: v,
	}
	q.do(0, len(v)-1)
}
