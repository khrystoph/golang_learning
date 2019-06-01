//sort algorithms in go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"
)

func init() {
	flag.IntVar(&arraysize, "arraysize", 10, "enter an integer")
	flag.Int64Var(&maxIntSize, "max", 1000, "enter an integer")
	flag.BoolVar(&printArray, "p", false, "tells the program to print the arrays at the end, but before the times.")
}

var (
	arraysize  int
	maxIntSize int64
	printArray bool
)

//NADA is used in mergesort as a sanity check to check if we should perform specific actions during sorting
const NADA int64 = -1

//THREADCOUNT is a constant to make adjusting the number of threads in waitgroup easier
const THREADCOUNT = 5

//MINLEN is the minimum length needed for goroutines to be effective. If we hit this value, we want to call regular mergesort
//instead of another go routine so we can maximize our efficiency
const MINLEN = 1000

//DeepCopy is a helper function for mergesort to copy the slices created by recursion to a new slice
func DeepCopy(vals []int64) []int64 {
	tmp := make([]int64, len(vals))
	copy(tmp, vals)
	return tmp
}

//mergesort recursively divides each array in half and sorts the smallest sizes, then merges them back after returning.
func mergesort(uintarray []int64) {

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
func tmergesort(uintarray []int64, r chan []int64) {
	if len(uintarray) < MINLEN {
		mergesort(uintarray)
		r <- uintarray
		return
	}

	leftchan := make(chan []int64)
	rightchan := make(chan []int64)
	middle := len(uintarray) / 2

	go tmergesort(uintarray[:middle], leftchan)
	go tmergesort(uintarray[middle:], rightchan)

	luintarray := <-leftchan
	ruintarray := <-rightchan

	r <- merge(luintarray, ruintarray)
	close(leftchan)
	close(rightchan)
	return

}

//merge is a function that actually performs the merge of two arrays for the mergesort algorithm.
func merge(left, right []int64) []int64 {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int64, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

//quicksort algorithm uses pivot value for divide and conquer to sort smaller arrays and put back together sorted
func quicksort(uintarray []int64) {

	if len(uintarray) > 1 {
		pivotIndex := len(uintarray) / 2
		var smallerItems = []int64{}
		var largerItems = []int64{}

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

		var merged = append(append(append([]int64{}, smallerItems...), []int64{uintarray[pivotIndex]}...), largerItems...)

		for j := 0; j < len(uintarray); j++ {
			uintarray[j] = merged[j]
		}

	}
	return
}

//heapify is a helper function for heapsort
func heapify(items []int64, idx int64, size int64) {
	l := 2*idx + 1 // left = 2*i + 1
	r := 2*idx + 2 // right = 2*i + 2

	var largest int64
	if l <= size && items[l] > items[idx] {
		largest = l
	} else {
		largest = idx
	}

	if r <= size && items[r] > items[largest] {
		largest = r
	}

	if largest != idx {
		t := items[idx]
		items[idx] = items[largest]
		items[largest] = t
		heapify(items, largest, size)
	}
	return
}

//heapsort function creates a heap of unsorted items and
func heapsort(items []int64) {
	var L int64
	L = int64(len(items)) //heap size
	m := int64(L / 2)     //middle

	for i := m; i >= 0; i-- {

		heapify(items, i, L-1)
	}

	F := L - 1
	for j := F; j >= 0; j-- {
		t := items[0]
		items[0] = items[j]
		items[j] = t
		F--
		heapify(items, 0, F)
	}
	return
}

//Bubblesort. See Readme for more info.
func bubblesort(uintarray []int64) {
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

//builtInSort is leveraging the package sort to do an efficient and stable sort
func builtInSort(uintarray []int64) {
	sort.Slice(uintarray, func(i, j int) bool {
		return uintarray[i] < uintarray[j]
	})
	return
}

func arrayprinter(sliceint []int64, arrayname string) (err error) {
	fmt.Printf("\nPrinting values for %s:\n", arrayname)
	for i := range sliceint {
		fmt.Printf("Value in %s[%d] is %v\n", arrayname, i, sliceint[i])
	}
	fmt.Printf("The address of pointer slicepointer is %p\n", sliceint)
	return nil
}

func createArray(arraysize int, maxSize int64) (intarray []int64) {
	for i := 0; i < arraysize; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val := r.Int63n(maxSize)
		intarray = append(intarray, val)
	}
	return intarray
}

func routineTimer(start time.Time, delta *time.Duration) {
	*delta = time.Since(start)
	return
}

func main() {
	var (
		bubbleint, mergeint, quickint, heapint, mergeintthread, builtInInt                                []int64
		bubblesortTimer, builtInSortTimer, quicksortTimer, mergesortTimer, tmergesortTimer, heapsortTimer time.Duration
		bswg, biwg, qswg, mswg, hswg                                                                      sync.WaitGroup
	)
	flag.Parse()
	runtime.GOMAXPROCS(4)

	intarray := createArray(arraysize, maxIntSize)

	bubbleint = append(bubbleint, intarray...)
	mergeint = append(mergeint, intarray...)
	quickint = append(quickint, intarray...)
	heapint = append(heapint, intarray...)
	mergeintthread = append(mergeintthread, intarray...)
	builtInInt = append(builtInInt, intarray...)

	bswg.Add(1)
	biwg.Add(1)
	qswg.Add(1)
	mswg.Add(1)
	hswg.Add(1)
	go func(somearray []int64) {
		start := time.Now()
		defer routineTimer(start, &bubblesortTimer)
		defer bswg.Done()
		bubblesort(somearray)
		fmt.Println("Finished Bubblesort.")
	}(bubbleint)
	go func(somearray []int64) {
		start := time.Now()
		defer routineTimer(start, &builtInSortTimer)
		defer biwg.Done()
		builtInSort(somearray)
		fmt.Println("Finish Built-in sort.")
	}(builtInInt)
	go func(somearray []int64) {
		start := time.Now()
		defer routineTimer(start, &quicksortTimer)
		defer qswg.Done()
		quicksort(somearray)
		fmt.Println("Finished Quicksort.")
	}(quickint)
	go func(somearray []int64) {
		start := time.Now()
		defer routineTimer(start, &mergesortTimer)
		defer mswg.Done()
		mergesort(somearray)
		fmt.Println("Finished mergesort.")
	}(mergeint)
	go func(somearray []int64) {
		start := time.Now()
		defer routineTimer(start, &heapsortTimer)
		defer hswg.Done()
		heapsort(somearray)
		fmt.Println("Finished heapsort.")
	}(heapint)
	bswg.Wait()
	biwg.Wait()
	qswg.Wait()
	mswg.Wait()
	hswg.Wait()

	testmergestring := make(chan string)
	rchan := make(chan []int64)
	var testmergearray []int64
	go func(timer *time.Duration) {
		defer routineTimer(time.Now(), timer)
		tmergesort(mergeintthread, rchan)
		testmergestring <- "Finished with Threaded Merge Sort."
	}(&tmergesortTimer)
	testmergearray = <-rchan
	fmt.Println(<-testmergestring)
	close(testmergestring)

	if printArray {
		arrayprinter(testmergearray, "mergeintthread")
	}

	fmt.Println("Bubblesort finished after: ", bubblesortTimer)
	fmt.Println("Quicksort finished after: ", quicksortTimer)
	fmt.Println("Heapsort finished after: ", heapsortTimer)
	fmt.Println("Mergesort finished after: ", mergesortTimer)
	fmt.Println("Built-in sort finished after: ", builtInSortTimer)
	fmt.Println("Threaded Mergesort finished after: ", tmergesortTimer)

	close(rchan)
}
