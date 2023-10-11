package insertion

// Time complexity: O(n^2)
// Memory complexity: O(1)
func SortInt(data []int) {
	for i, curVal := range data {
		for ; (i > 0) && curVal < data[i-1]; i-- {
			data[i] = data[i-1]
		}
		data[i] = curVal
	}
}
