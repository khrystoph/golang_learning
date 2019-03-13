//sort algorithms in go

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func mergesort(uintarray []int) (err error) {
	fmt.Printf("entering mergesort.")
	return nil
}

func quicksort(uintarray []int) (quicksorted *[]int, err error) {
	fmt.Printf("entering quicksort.")
	return nil, nil
}

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
	err = nil

	return err
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
