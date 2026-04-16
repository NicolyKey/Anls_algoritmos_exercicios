package main

import (
	"fmt"
	"math/rand"
	"time"
)

func benchInsertRandomSlice() time.Duration {
	s := make([]int, 0)
	start := time.Now()
	for i := 0; i < N; i++ {
		pos := 0
		if len(s) > 0 {
			pos = rand.Intn(len(s) + 1)
		}
		sliceInsertAt(&s, pos, i)
	}
	return time.Since(start)
}

func benchInsertRandomLinked() time.Duration {
	ll := NewLinkedList()
	start := time.Now()
	for i := 0; i < N; i++ {
		pos := 0
		if ll.Size > 0 {
			pos = rand.Intn(ll.Size + 1)
		}
		ll.InsertAt(pos, i)
	}
	return time.Since(start)
}

func RunQuestao2() {
	fmt.Println("========================================")
	fmt.Println("  QUESTAO 2 - Insercao em posicao aleatoria")
	fmt.Println("========================================")

	runTests("Q2 - Insercao em posicao aleatoria", 20, benchInsertRandomSlice, benchInsertRandomLinked)
}
