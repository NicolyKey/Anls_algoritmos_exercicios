package main

import (
	"fmt"
	"time"
)

func benchRemoveFirstSlice() time.Duration {
	s := buildSlice()
	start := time.Now()
	for len(s) > 0 {
		sliceRemoveFirst(&s)
	}
	return time.Since(start)
}

func benchRemoveFirstLinked() time.Duration {
	ll := buildLinkedList()
	start := time.Now()
	for ll.Size > 0 {
		ll.RemoveFirst()
	}
	return time.Since(start)
}

func benchRemoveLastSlice() time.Duration {
	s := buildSlice()
	start := time.Now()
	for len(s) > 0 {
		sliceRemoveLast(&s)
	}
	return time.Since(start)
}

func benchRemoveLastLinked() time.Duration {
	ll := buildLinkedList()
	start := time.Now()
	for ll.Size > 0 {
		ll.RemoveLast()
	}
	return time.Since(start)
}

func RunQuestao3() {
	fmt.Println("========================================")
	fmt.Println("  QUESTAO 3 - Remocao de elementos")
	fmt.Println("========================================")

	runTests("Q3.1 - Remocao do primeiro elemento", 20, benchRemoveFirstSlice, benchRemoveFirstLinked)

	runTests("Q3.2 - Remocao do ultimo elemento", 20, benchRemoveLastSlice, benchRemoveLastLinked)
}
