package insertion

func SortInt(data []int) {
	for i, curVal := range data {
		for ; (i > 0) && curVal < data[i-1]; i-- {
			data[i] = data[i-1]
		}
		data[i] = curVal
	}
}
