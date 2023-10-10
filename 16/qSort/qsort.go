package qsort

import "github.com/Ckakalka/wbLevel1/16/insertion"

// This constant should be selected based on the results of profiling
// In our case, the value is taken from the ceiling
const insertionBound = 16

// Time complexity: average O(nlog(n)), worst O(n^2)
// Memory complexity (for stack): average O(log(n)), worst O(n).
func SortInt(arr []int) {
	if len(arr) <= insertionBound {
		insertion.SortInt(arr)
		return
	}
	pivot := arr[len(arr)/2]
	left := 0
	right := len(arr) - 1
	for left < right {
		for left < right && arr[left] < pivot {
			left++
		}
		for left < right && arr[right] > pivot {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	if left != right {
		right++
	}
	SortInt(arr[:right])
	SortInt(arr[left:])
}
