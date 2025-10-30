package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		len         int
		expectedLen int
	}{
		{
			name:        "zero size",
			len:         0,
			expectedLen: 0,
		},
		{
			name:        "non-negative sizwe",
			len:         5,
			expectedLen: 5,
		},
	}

	for _, test := range tests {
		require.Len(t, generateRandomElements(test.len), test.expectedLen)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected int
	}{
		{
			name:     "three non-negative values",
			slice:    []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "only zeros",
			slice:    []int{0, 0},
			expected: 0,
		},
		{
			name:     "one unit",
			slice:    []int{1},
			expected: 1,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			expected: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maximum(test.slice))
	}
}
