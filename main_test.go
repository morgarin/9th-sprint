package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	expLen := 5
	require.Equal(t, len(generateRandomElements(expLen)), expLen)
}
func TestMaximum(t *testing.T) {
	slice := []int{0, 0}
	assert.Equal(t, maximum(slice), 0)
	slice = []int{1}
	assert.Equal(t, maximum(slice), 1)
	slice = []int{1, 2, 3}
	assert.Equal(t, maximum(slice), 3)
	slice = []int{}
	assert.Equal(t, maximum(slice), 0)
}
