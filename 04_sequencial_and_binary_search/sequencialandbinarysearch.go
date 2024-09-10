package searchandbinarysearch

type Collection struct {
	values []int
}

func NewCollection(values []int) *Collection {
	return &Collection{
		values: values,
	}
}

func (c Collection) IndexOfSequencial(value int) int {
	for idx, v := range c.values {
		if v == value {
			return idx
		}
	}
	return -1
}

func (c Collection) IndexOfBinary(value int) int {
	start := 0
	end := len(c.values) - 1
	for start <= end {
		mid := (start + end) / 2
		v := c.values[mid]
		if v == value {
			return mid
		}
		if value < v {
			end = mid - 1
		} else if value > v {
			start = mid + 1
		}
	}
	return -1
}
