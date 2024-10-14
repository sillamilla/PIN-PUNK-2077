package service

import (
	"MiniGame-PinUp/MatrixSequence_Service/internal/models"
	"math/rand"
)

type Matrix interface {
	NewMetricData() models.MatrixData
}

type service struct {
}

func NewService() Matrix {
	return &service{}
}

func (s *service) NewMetricData() models.MatrixData {
	matrix := createMetric()
	keySequence := newKeySequence()

	return models.MatrixData{
		Matrix:      matrix,
		KeySequence: keySequence,
	}
}

func createMetric() models.Matrix {
	var metric models.Matrix

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			metric.Data[i][j] = rand.Intn(10)
		}
	}

	return metric
}

func newKeySequence() models.KeySequence {
	var keySequence models.KeySequence

	for i := 0; i < 7; i++ {
		key := rand.Intn(10)
		keySequence.Keys = append(keySequence.Keys, key)
	}

	return keySequence
}
