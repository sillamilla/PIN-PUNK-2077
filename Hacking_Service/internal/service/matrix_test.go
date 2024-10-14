package service

import (
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"testing"
)

func TestCountMatches(t *testing.T) {
	service := NewService(nil)

	tests := []struct {
		name          string
		matrix        models.Matrix
		keySequence   models.KeySequence
		expectedKeys  string
		expectedCount int
	}{
		{
			name: "Test with 1 match",
			matrix: models.Matrix{
				Data: [5][5]int{
					{1, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			keySequence:   models.KeySequence{Keys: []int{2, 9, 1, 5}},
			expectedKeys:  "1",
			expectedCount: 0,
		},
		{
			name: "Test with 3 matches",
			matrix: models.Matrix{
				Data: [5][5]int{
					{1, 0, 0, 0, 0},
					{0, 4, 5, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{2, 0, 0, 0, 3},
				},
			},
			keySequence:   models.KeySequence{Keys: []int{1, 2, 3}},
			expectedKeys:  "123",
			expectedCount: 1,
		},
		{
			name: "Test with 4 matches",
			matrix: models.Matrix{
				Data: [5][5]int{
					{1, 9, 0, 0, 0},
					{2, 5, 6, 3, 0},
					{0, 0, 0, 10, 0},
					{0, 5, 0, 0, 0},
					{0, 0, 0, 4, 0},
				},
			},
			keySequence:   models.KeySequence{Keys: []int{1, 2, 3, 4}},
			expectedKeys:  "1234",
			expectedCount: 2,
		},
		{
			name: "Test with 4 matches",
			matrix: models.Matrix{
				Data: [5][5]int{
					{0, 1, 0, 0, 0},
					{0, 0, 7, 0, 6},
					{0, 2, 3, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 4, 0, 5},
				},
			},
			keySequence:   models.KeySequence{Keys: []int{1, 2, 3, 4, 5, 6, 7}},
			expectedKeys:  "1234567",
			expectedCount: 3,
		},
		{
			name: "Test with 6 matches",
			matrix: models.Matrix{
				Data: [5][5]int{
					{0, 1, 0, 1, 1}, //die
					{0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 1, 1, 0},
				},
			},
			keySequence:   models.KeySequence{Keys: []int{1, 4, 9}},
			expectedKeys:  "1111111",
			expectedCount: 3,
		},
		{
			name: "Test with no matches",
			matrix: models.Matrix{
				Data: [5][5]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			keySequence:   models.KeySequence{Keys: []int{8, 9}},
			expectedKeys:  "",
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, count := service.CountMatches(tt.matrix, tt.keySequence)

			if keys != tt.expectedKeys {
				t.Errorf("expected keys %s, got %s", tt.expectedKeys, keys)
			}
			if count != tt.expectedCount {
				t.Errorf("expected count %d, got %d", tt.expectedCount, count)
			}
		})
	}
}
