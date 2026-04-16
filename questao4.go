package main

import (
	"fmt"
	"math/rand"
	"time"
)

func benchRemoveRandomSlice() time.Duration {
	s := buildSlice()
	start := time.Now()
	for len(s) > 0 {
		pos := rand.Intn(len(s))
		sliceRemoveAt(&s, pos)
	}
	return time.Since(start)
}

func benchRemoveRandomLinked() time.Duration {
	ll := buildLinkedList()
	start := time.Now()
	for ll.Size > 0 {
		pos := rand.Intn(ll.Size)
		ll.RemoveAt(pos)
	}
	return time.Since(start)
}

func RunQuestao4() {
	fmt.Println("========================================")
	fmt.Println("  QUESTAO 4 - Remocao em indice aleatorio")
	fmt.Println("========================================")

	runTests("Q4 - Remocao em indice aleatorio", 20, benchRemoveRandomSlice, benchRemoveRandomLinked)
}
