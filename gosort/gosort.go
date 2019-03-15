//sort algorithms in go

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

/*stub out single-threaded mergesort algorithm's function call.
Expected big O time is O(nlog n) in average and worst case.
Expectation is that this will be the second fastest to complete on average.
However, this should come out 2-X times slower than threaded mergesort as
the threads should have a speed multiplier based on the number of CPUs
the go scheduler can/will schedule the threads across.
*/
func mergesort(uintarray []int) (err error) {
	fmt.Printf("entering mergesort.")
	return nil
}

/*stub out threaded mergesort algorithms's function call.
This algorithm is expected to get variable performance, but better than nlog n.
I expect that with a 4-core processor, we should see about a 4x gain in speed.
Likely it will be a bit slower than 4x due to context switching and memory access
patterns, but there will be tests run on several sizes of processor (including 20+ cores).
Once everything works as expected.
*/
func tmergesort(uintarray []int) (err error) {
	fmt.Printf("entering threaded mergesort.")
	return nil
}

//stubbing out quicksort algorithm's function call
func quicksort(uintarray []int) (err error) {
	fmt.Printf("entering quicksort.")
	return nil
}

//stubbing out heapsort algorithm's function call
func heapsort(uintarray []int) (err error) {
	fmt.Printf("entering heapsort")
	return nil
}

/*
Bubble sort algorithm. This algorithm was the original and slow. results should
always be slower than the other algorithms as expected O is O(n2).
*/
func bubblesort(uintarray []int) (err error) {
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

	return nil
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
		intarray, bubbleint, mergeint, quickint, heapint, mergeintthread []int
		arraysize                                                        int
		val                                                              int
		err                                                              error
	)

	flag.IntVar(&arraysize, "arraysize", 10, "enter an integer")

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

	if bubblesort(bubbleint); err != nil {
		fmt.Printf("Error sorting with bubble sort. Error msg: %v\n", err)
	}

	fmt.Printf("Pointers to arrays:\nintarray:%p\nbubbleint:%p\nmergeint:%p\n"+
		"quickint:%p\nheapint:%p\nmergeintthread:%p", intarray, bubbleint, mergeint,
		quickint, heapint, mergeintthread)

	arrayprinter(intarray, "intarray")
	arrayprinter(bubbleint, "bubbleint")
	arrayprinter(mergeint, "mergeint")
	arrayprinter(quickint, "quickint")
	arrayprinter(mergeintthread, "mergeintthread")

	return
}
