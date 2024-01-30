package main

import (
	"testing"
)

func TestDataRaceCondition(t *testing.T) {

}

// Usando Atomic Operations para evitar race conditions
// func TestDataRaceCondition(t *testing.T) {
// 	var state int32

// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			// state += int32(i)
// 			atomic.AddInt32(&state, int32(i))
// 		}(i)
// 	}
// }

// evitando race conditions com mutex
// PS ---> pode ter um alto custo computacional e deixar o sistema lento
// func TestDataRaceCondition(t *testing.T) {
// 	var state int32
// 	var mu sync.RWMutex

// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			mu.Lock()
// 			state += int32(i)
// 			// business logic
// 			mu.Unlock()
// 		}(i)
// 	}
// }
