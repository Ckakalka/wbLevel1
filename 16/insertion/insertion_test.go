package insertion

import (
	"testing"
)

var dataToSort = [...][]int{
	{5, 2, 3, 3, 1, -10},
	{5},
	{},
}

var expectedResults = [...][]int{
	{-10, 1, 2, 3, 3, 5},
	{5},
	{},
}

func TestSortInt(t *testing.T) {
	for i, d := range dataToSort {
		data := make([]int, len(d))
		copy(data, d)
		SortInt(data)
		for j, v := range data {
			if v != expectedResults[i][j] {
				t.Errorf("Test set %d. Expected value %d but value %d received\n",
					i+1, expectedResults[i][j], v)
			}
		}
	}
}
