package main

import (
	"fmt"
	"time"
)

func benchInsertEndSlice() time.Duration {
	start := time.Now()
	s := make([]int, 0)
	for i := 0; i < N; i++ {
		s = append(s, i)
	}
	return time.Since(start)
}

func benchInsertEndSliceWithCap(cap int) time.Duration {
	start := time.Now()
	s := make([]int, 0, cap)
	for i := 0; i < N; i++ {
		s = append(s, i)
	}
	return time.Since(start)
}

func benchInsertEndLinked() time.Duration {
	start := time.Now()
	ll := NewLinkedList()
	for i := 0; i < N; i++ {
		ll.InsertEnd(i)
	}
	return time.Since(start)
}

func RunQuestao1() {
	fmt.Println("========================================")
	fmt.Println("  QUESTAO 1 - Insercao ao final")
	fmt.Println("========================================")

	runTests("Q1.1 - Insercao no final", 20, benchInsertEndSlice, benchInsertEndLinked)

	fmt.Println("=== Q1.2 - Influencia da capacidade inicial (Slice) ===")
	caps := []int{10, 1000, 100_000}
	for _, c := range caps {
		times := make([]time.Duration, 10)
		for i := 0; i < 10; i++ {
			times[i] = benchInsertEndSliceWithCap(c)
		}
		avg := avgDuration(times)
		fmt.Printf("  Capacidade inicial %d: media = %v\n", c, avg)
	}
	fmt.Println()
}
