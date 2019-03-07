//sort algorithms in go

package gosort

import (
	"fmt"
)

func mergesort(uintarray []int) (err error) {
	fmt.Printf("entering mergesort.")
	return nil
}

func quicksort(uintarray []int) (err error, quicksorted *[]int) {
	return nil, nil
}

func bubblesort(uintarray []int) (bubblesorted *[]int, err error) {
	var (
		tempA, i, j int
	)

	for i = 0; i < len(uintarray)-1; i++ {
		for j = 0; j < len(uintarray)-i-1; j++ {
			if uintarray[j] > uintarray[j+1] {
				tempA = uintarray[j]
				uintarray[j] = uintarray[j+1]
				uintarray[j+i] = tempA
			}
		}
	}
	return nil, nil
}

func arrayprinter(sortedarray []int) (err error) {
	return
}

func main() {
	return
}
