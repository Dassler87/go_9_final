package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {

	// Проверка, что при запросе 0 элементов возвращается пустой срез.
	t.Run("zero_size", func(t *testing.T) {
		data := generateRandomElements(0)
		if len(data) != 0 {
			t.Error("Expected empty slice for size 0")
		}
	})

	// Проверка, что при отрицательном размере возвращается пустой срез.
	t.Run("negative_size", func(t *testing.T) {
		data := generateRandomElements(-10)
		if len(data) != 0 {
			t.Error("Expected empty slice for negative size")
		}
	})

	// Проверка, что генерируется срез нужной длины и все числа попадают
	// в заданный диапазон (от 1 до 100000)
	t.Run("positive_size_and_range", func(t *testing.T) {
		size := 1000
		data := generateRandomElements(size)
		if len(data) != size {
			t.Errorf("Expected length %d, got %d", size, len(data))
		}
		for _, v := range data {
			if v <= 0 || v > 100000 {
				t.Errorf("Value %d out of expected range (1..100000)", v)
			}
		}
	})
}

func TestMaximum(t *testing.T) {

	// Проверка базового случая для пустой коллекции.
	t.Run("empty_slice", func(t *testing.T) {
		val := maximum([]int{})
		if val != 0 {
			t.Error("Expected 0 for empty slice")
		}
	})

	// Проверка работы с одним элементом.
	t.Run("single_element", func(t *testing.T) {
		val := maximum([]int{100})
		if val != 100 {
			t.Errorf("Expected 100, got %d", val)
		}
	})

	// Проверка поиска максимума в обычном массиве.
	t.Run("multiple_elements", func(t *testing.T) {
		val := maximum([]int{3, 9, 1, 7, 2})
		if val != 9 {
			t.Errorf("Expected 9, got %d", val)
		}
	})

	// Проверка, что функция не ломается, если все элементы равны.
	t.Run("all_equal", func(t *testing.T) {
		val := maximum([]int{5, 5, 5})
		if val != 5 {
			t.Errorf("Expected 5, got %d", val)
		}
	})
}

func TestMaxChunks(t *testing.T) {

	// Базовая проверка для параллельной версии.
	t.Run("empty_slice", func(t *testing.T) {
		val := maxChunks([]int{})
		if val != 0 {
			t.Error("Expected 0 for empty slice")
		}
	})

	// Проверка для одного элемента.
	t.Run("single_element", func(t *testing.T) {
		val := maxChunks([]int{123})
		if val != 123 {
			t.Errorf("Expected 123, got %d", val)
		}
	})

	// Генерирует большой массив и сравнивает результат параллельной функции maxChunks
	// с результатом последовательной maximum
	t.Run("compare_with_single_thread", func(t *testing.T) {
		data := generateRandomElements(100_000)
		single := maximum(data)
		multi := maxChunks(data)
		if single != multi {
			t.Errorf("Results mismatch: single=%d, multi=%d", single, multi)
		}
	})
}
