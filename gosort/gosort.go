//sort algorithms in go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	flag.IntVar(&arraysize, "arraysize", 10, "enter an integer")
}

var (
	arraysize int
)

//NADA is used in mergesort as a sanity check to check if we should perform specific actions during sorting
const NADA int = -1

//DeepCopy is a helper function for mergesort to copy the slices created by recursion to a new slice
func DeepCopy(vals []int) []int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	return tmp
}

//mergesort recursively divides each array in half and sorts the smallest sizes, then merges them back after returning.
func mergesort(uintarray []int) {

	if len(uintarray) > 1 {
		mid := len(uintarray) / 2
		left := DeepCopy(uintarray[0:mid])
		right := DeepCopy(uintarray[mid:])

		mergesort(left)
		mergesort(right)

		l := 0
		r := 0

		for i := 0; i < len(uintarray); i++ {

			lval := NADA
			rval := NADA

			if l < len(left) {
				lval = left[l]
			}

			if r < len(right) {
				rval = right[r]
			}

			if (lval != NADA && rval != NADA && lval < rval) || rval == NADA {
				uintarray[i] = lval
				l++
			} else if (lval != NADA && rval != NADA && lval >= rval) || lval == NADA {
				uintarray[i] = rval
				r++
			}

		}
	}

}

//stub out threaded mergesort algorithms's function call.
//See the readme for more info on tmergesort.
func tmergesort(uintarray []int) (err error) {
	fmt.Printf("entering threaded mergesort.")
	return nil
}

//quicksort algorithm uses pivot value for divide and conquer to sort smaller arrays and put back together sorted
func quicksort(uintarray []int) {

	if len(uintarray) > 1 {
		pivotIndex := len(uintarray) / 2
		var smallerItems = []int{}
		var largerItems = []int{}

		for i := range uintarray {
			val := uintarray[i]
			if i != pivotIndex {
				if val < uintarray[pivotIndex] {
					smallerItems = append(smallerItems, val)
				} else {
					largerItems = append(largerItems, val)
				}
			}
		}

		quicksort(smallerItems)
		quicksort(largerItems)

		var merged = append(append(append([]int{}, smallerItems...), []int{uintarray[pivotIndex]}...), largerItems...)

		for j := 0; j < len(uintarray); j++ {
			uintarray[j] = merged[j]
		}

	}

}

//stubbing out heapsort algorithm's function call
func heapsort(uintarray []int) (err error) {
	fmt.Printf("entering heapsort")
	return nil
}

//Bubblesort. See Readme for more info.
func bubblesort(uintarray []int) {
	var (
		intlen  = len(uintarray)
		swapped = true
	)
	for swapped {
		swapped = false
		for i := 0; i < (intlen - 1); i++ {
			if uintarray[i] > uintarray[i+1] {
				uintarray[i], uintarray[i+1] = uintarray[i+1], uintarray[i]
				swapped = true
			}
		}
	}

	return
}

func builtInSort(uintarray []int) {
	sort.Slice(uintarray, func(i, j int) bool {
		return uintarray[i] < uintarray[j]
	})
	return
}

func arrayprinter(sliceint []int, arrayname string) (err error) {
	fmt.Printf("\nPrinting values for %s:\n", arrayname)
	for i := range sliceint {
		fmt.Printf("Value in %s[%d] is %v\n", arrayname, i, sliceint[i])
	}
	fmt.Printf("The address of pointer slicepointer is %p\n", sliceint)
	return nil
}

func main() {
	var (
		intarray, bubbleint, mergeint, quickint, heapint, mergeintthread, builtInInt []int
		val                                                                          int
	)

	flag.Parse()

	for i := 0; i < arraysize; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val = r.Intn(1000)
		intarray = append(intarray, val)
	}

	bubbleint = append(bubbleint, intarray...)
	mergeint = append(mergeint, intarray...)
	quickint = append(quickint, intarray...)
	heapint = append(heapint, intarray...)
	mergeintthread = append(mergeintthread, intarray...)
	builtInInt = append(builtInInt, intarray...)

	bubblesort(bubbleint)
	builtInSort(builtInInt)
	quicksort(quickint)
	mergesort(mergeint)

	fmt.Printf("Pointers to arrays:\nintarray:%p\nbubbleint:%p\nmergeint:%p\n"+
		"quickint:%p\nheapint:%p\nmergeintthread:%p", intarray, bubbleint, mergeint,
		quickint, heapint, mergeintthread)

	arrayprinter(intarray, "intarray")
	arrayprinter(bubbleint, "bubbleint")
	arrayprinter(builtInInt, "builtInInt")
	arrayprinter(mergeint, "mergeint")
	arrayprinter(quickint, "quickint")
	arrayprinter(mergeintthread, "mergeintthread")

	return
}
