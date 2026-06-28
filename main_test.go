package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantNil bool
		wantLen int
	}{
		{name: "size_zero", size: 0, wantNil: true, wantLen: 0},
		{name: "size_negative", size: -5, wantNil: true, wantLen: 0},
		{name: "size_one", size: 1, wantNil: false, wantLen: 1},
		{name: "size_small", size: 10, wantNil: false, wantLen: 10},
		{name: "size_large", size: 100_000, wantNil: false, wantLen: 100_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаём генератор с фиксированным seed — чтобы тест был воспроизводимым
			r := rand.New(rand.NewSource(42))

			data := generateRandomElements(tt.size, r)

			if tt.wantNil {
				assert.Nil(t, data, "Ожидался nil для size=%d", tt.size)
				return
			}

			require.NotNil(t, data, "Для положительного размера не должен возвращаться nil")
			assert.Len(t, data, tt.wantLen, "Длина слайса должна быть %d, а была %d", tt.wantLen, len(data))

			// Теперь проверяем только неотрицательность: r.Int() всегда >= 0
			for i, v := range data {
				assert.True(t, v >= 0,
					"data[%d] = %d — отрицательное значение, а r.Int() не должен его возвращать", i, v)

				// Опционально: можно убедиться, что значения не превышают максимально возможное int
				assert.True(t, v <= math.MaxInt,
					"data[%d] = %d превышает math.MaxInt", i, v)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{name: "empty_slice", input: []int{}, expect: 0},
		{name: "single_element", input: []int{42}, expect: 42},
		{name: "multiple_elements", input: []int{3, 9, 1, 7, 2}, expect: 9},
		{name: "all_equal", input: []int{5, 5, 5}, expect: 5},
		{name: "negative_and_positive", input: []int{-10, -3, 0, 7, 2}, expect: 7},
		{name: "max_at_start", input: []int{100, 10, 20}, expect: 100},
		{name: "max_at_end", input: []int{10, 20, 100}, expect: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.input)
			assert.Equal(t, tt.expect, got,
				"maximum(%v) = %d; ожидалось %d", tt.input, got, tt.expect)
		})
	}
}
