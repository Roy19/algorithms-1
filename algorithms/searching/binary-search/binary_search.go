package bs

import (
	"golang.org/x/exp/constraints"
)

func search[T constraints.Ordered](sortedArray []T, el T) int {
	init, end := 0, len(sortedArray)-1

	for init <= end {
		middle := ((end - init) >> 1) + init

		if sortedArray[middle] == el {
			return middle
		}

		if sortedArray[middle] < el {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}
