package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size == 0 {
		return []int{}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := range slice {
		slice[i] = r.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	switch {
	case len(data) == 0:
		return 0
	case len(data) == 1:
		return data[0]
	default:
		max := 0
		for i := range data {
			if data[i] > max {
				max = data[i]
			}
		}
		return max
	}
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	partSize := len(data) / CHUNKS
	finalMax := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * partSize
		end := start + partSize
		if i == (CHUNKS - -1) {
			end = len(data)
		}
		go func(data []int) {
			defer wg.Done()
			finalMax[i] = maximum(data)
		}(data[start:end])
	}
	wg.Wait()

	return maximum(finalMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
