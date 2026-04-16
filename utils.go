package main

import (
	"fmt"
	"time"
)

const N = 100_000

func avgDuration(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func printComparison(avgSlice, avgLinked time.Duration) {
	fmt.Printf("  Slice (vetor):     %v\n", avgSlice)
	fmt.Printf("  LinkedList:        %v\n", avgLinked)
	if avgSlice < avgLinked {
		ratio := float64(avgLinked) / float64(avgSlice)
		fmt.Printf("  -> Slice foi %.2fx mais rapida\n", ratio)
	} else {
		ratio := float64(avgSlice) / float64(avgLinked)
		fmt.Printf("  -> LinkedList foi %.2fx mais rapida\n", ratio)
	}
	fmt.Println()
}

func runTests(name string, runs int, sliceFn, linkedFn func() time.Duration) {
	sliceTimes := make([]time.Duration, runs)
	linkedTimes := make([]time.Duration, runs)

	for i := 0; i < runs; i++ {
		sliceTimes[i] = sliceFn()
		linkedTimes[i] = linkedFn()
	}

	avgS := avgDuration(sliceTimes)
	avgL := avgDuration(linkedTimes)

	fmt.Printf("=== %s (%d execucoes) ===\n", name, runs)
	printComparison(avgS, avgL)
}

func buildSlice() []int {
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = i
	}
	return s
}

func buildLinkedList() *LinkedList {
	ll := NewLinkedList()
	for i := 0; i < N; i++ {
		ll.InsertEnd(i)
	}
	return ll
}

func sliceInsertAt(s *[]int, index, value int) {
	*s = append(*s, 0)
	copy((*s)[index+1:], (*s)[index:])
	(*s)[index] = value
}

func sliceRemoveFirst(s *[]int) {
	*s = (*s)[1:]
}

func sliceRemoveLast(s *[]int) {
	*s = (*s)[:len(*s)-1]
}

func sliceRemoveAt(s *[]int, index int) {
	*s = append((*s)[:index], (*s)[index+1:]...)
}
