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

func quicksort(uintarray []int) (quicksorted *[]int, interr error) {
	return nil, nil
}

func bubblesort(uintarray *[]int) (bubblesorted *[]int, err error) {
	var (
		intlen   = len(*uintarray)
		intarray = *uintarray
		swapped  = true
	)
	for swapped {
		swapped = false
		for i := 0; i < (intlen - 1); i++ {
			if intarray[i] > intarray[i+1] {
				intarray[i], intarray[i+1] = intarray[i+1], intarray[i]
				swapped = true
			}
		}
	}

	bubblesorted = &intarray
	err = nil

	return bubblesorted, err
}

func arrayprinter(slicepntr *[]int) (err error) {
	var slicearray = *slicepntr
	for i := range slicearray {
		fmt.Printf("Value in slicepntr[%d] is %v\n", i, slicearray[i])
	}
	fmt.Printf("The address of pointer slicepointer is %p\n", slicepntr)
	return nil
}

func main() {
	var (
		intarray    = []int{}
		intarrayptr = &intarray
		arraysize   = 10
		val         int
		err         error
	)

	flag.IntVar(&arraysize, "arraysize", 10, "enter an integer")

	flag.Parse()

	for i := 0; i < arraysize; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val = r.Intn(1000)
		intarray = append(intarray, val)
	}

	if intarrayptr, err = bubblesort(intarrayptr); err != nil {
		fmt.Printf("Error sorting with bubble sort. Error msg: %v\n", err)
	}

	intarray = *intarrayptr
	arrayprinter(&intarray)

	return
}
