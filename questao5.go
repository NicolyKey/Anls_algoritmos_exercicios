package main

import (
	"fmt"
	"math/rand"
	"time"
)

func benchAccessSlice() time.Duration {
	s := buildSlice()
	start := time.Now()
	for i := 0; i < 10_000; i++ {
		_ = s[rand.Intn(len(s))]
	}
	return time.Since(start)
}

func benchAccessLinked() time.Duration {
	ll := buildLinkedList()
	start := time.Now()
	for i := 0; i < 10_000; i++ {
		ll.Get(rand.Intn(ll.Size))
	}
	return time.Since(start)
}

func RunQuestao5() {
	fmt.Println("========================================")
	fmt.Println("  QUESTAO 5 - Acesso aleatorio (10.000 acessos)")
	fmt.Println("========================================")

	runTests("Q5 - Acesso aleatorio (10.000 acessos)", 20, benchAccessSlice, benchAccessLinked)
}
