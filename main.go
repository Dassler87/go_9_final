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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	// Обработка крайнего случая: если размер <= 0, возвращаем пустой слайс
	if size <= 0 {
		return []int{}
	}

	// Создаём локальный генератор случайных чисел.
	r := rand.New(rand.NewSource(87))

	result := make([]int, size)
	for i := range result {
		// Генерируем число в диапазоне [1, 1000000)
		result[i] = r.Intn(1000000) + 1
	}
	return result

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}
	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	n := len(data)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return data[0]
	}

	// Вычисляем размер одного чанка с округлением вверх.
	chunkSize := (n + CHUNKS - 1) / CHUNKS
	localMaxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for c := 0; c < CHUNKS; c++ {
		go func(chunkID int) {
			defer wg.Done()

			start := chunkID * chunkSize
			end := start + chunkSize
			if end > n {
				end = n
			}

			// Защита от пустого диапазона. Если CHUNKS > len(data), некоторые горутины
			// могут получить пустой чанк [start:end), где start == end.
			// Попытка доступа к chunk[0] в этом случае привела бы к панике,
			// поэтому мы явно задаем максимум для такого чанка как 0.
			if start >= end {
				localMaxes[chunkID] = 0
				return
			}

			chunk := data[start:end]
			localMax := chunk[0]
			for i := 1; i < len(chunk); i++ {
				if chunk[i] > localMax {
					localMax = chunk[i]
				}
			}
			localMaxes[chunkID] = localMax
		}(c)
	}

	// Ожидаем, пока все горутины не закончат свою работу.
	wg.Wait()

	// Находим максимум среди локальных максимумов
	globalMax := localMaxes[0]
	for i := 1; i < CHUNKS; i++ {
		if localMaxes[i] > globalMax {
			globalMax = localMaxes[i]
		}
	}
	return globalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println()

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)
	fmt.Println()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed.Microseconds())
}
