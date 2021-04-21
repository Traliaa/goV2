package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	example_3()
}

func example_1() {
	wg := sync.WaitGroup{}
	worker := 10

	wg.Add(worker)

	for i := 0; i < worker; i += 1 {
		go func(i int) {
			fmt.Printf("runner %d", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//Реализуйте функцию для разблокировки мьютекса с помощью defer
func example_2() {

}

func example_3() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	mm := map[int]int{}
	worker := 1000

	wg.Add(worker)

	for i := 0; i < worker; i += 1 {
		go func(i int) {
			defer wg.Done()
			if rand.Float32() < 0.1 {
				mu.Lock()
				mm[i] = i * 2
				mu.Unlock()
			}
			mu.Lock()
			_ = mm[i]
			mu.Unlock()

		}(i)
	}
	wg.Wait()
}
func example_3RW() {
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	mm := map[int]int{}
	worker := 10000

	wg.Add(worker)

	for i := 0; i < worker; i += 1 {
		go func(i int) {
			defer wg.Done()
			if rand.Float32() < 0.1 {
				mu.Lock()
				mm[i] = i * 2
				mu.Unlock()
			}
			mu.RLock()
			_ = mm[i]
			mu.RUnlock()

		}(i)
	}
	wg.Wait()
}
