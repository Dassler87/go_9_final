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
func generateRandomElements(size int, r *rand.Rand) []int {

	// При некорректном размере возвращаем nil
	if size <= 0 {
		return nil
	}

	result := make([]int, size)
	for i := range result {
		// Неограниченная случайность: любое неотрицательное int
		result[i] = r.Int()
	}
	return result

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
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

		// Подготовка слайса до запуска горутины
		start := c * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}

		chunk := data[start:end]

		go func(id int, ch []int) {
			defer wg.Done()

			// Пустой диапазон: используем maximum для получения результата.
			localMaxes[id] = maximum(ch)
		}(c, chunk) // Передаем ID и готовый слайс
	}

	// Ожидаем, пока все горутины не закончат свою работу.
	wg.Wait()

	// Используем maximum для финального поиска максимума
	return maximum(localMaxes)
}

func main() {
	// Создаём генератор со случайным seed (на основе времени)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("Генерируем %d целых чисел", SIZE)

	data := generateRandomElements(SIZE, r)
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
