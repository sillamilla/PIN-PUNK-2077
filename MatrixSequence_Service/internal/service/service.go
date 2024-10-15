package service

import (
	"MiniGame-PinUp/MatrixSequence_Service/internal/models"
	"MiniGame-PinUp/MatrixSequence_Service/pkg/hackService"
	"context"
	"errors"
	"math/rand"
)

type MatrixService interface {
	NewMatrixData(ctx context.Context) *models.MatrixData
	HackMatrix(ctx context.Context) (string, error)
}

type service struct {
	hackServiceClient hackService.Client
	matrixData        *models.MatrixData
}

func New(hackServiceClient hackService.Client) MatrixService {
	return &service{hackServiceClient: hackServiceClient}
}

func (s *service) NewMatrixData(_ context.Context) *models.MatrixData {
	matrix := s.createMatrix()
	keySequence := s.newKeySequence()

	s.matrixData = &models.MatrixData{
		Matrix:      matrix,
		KeySequence: keySequence,
	}

	return s.matrixData
}

func (s *service) HackMatrix(ctx context.Context) (string, error) {
	if s.matrixData == nil {
		return "", errors.New("matrix data is empty") //todo models
	}

	hackServiceMatrix := matrixToHackServiceMatrix(s.matrixData)
	response, err := s.hackServiceClient.HackMatrix(ctx, hackServiceMatrix)
	if err != nil {
		return "", err
	}

	return response.Status, nil
}

func matrixToHackServiceMatrix(matrix *models.MatrixData) hackService.MatrixData {
	var hackServiceMatrix hackService.MatrixData

	for i := range matrix.Data {
		for j := range matrix.Data[i] {
			hackServiceMatrix.Data[i][j] = matrix.Data[i][j]
		}
	}

	hackServiceMatrix.KeySequence.Keys = make([]int, len(matrix.KeySequence.Keys))
	copy(hackServiceMatrix.KeySequence.Keys, matrix.KeySequence.Keys)

	return hackServiceMatrix
}

func (s *service) createMatrix() models.Matrix {
	var matrix models.Matrix

	for i := 0; i < models.MatrixWidth; i++ {
		for j := 0; j < models.MatrixHeight; j++ {
			matrix.Data[i][j] = rand.Intn(10)
		}
	}

	return matrix
}

func (s *service) newKeySequence() models.KeySequence {
	var keySequence models.KeySequence

	for i := 0; i < models.KeySequenceLength; i++ {
		key := rand.Intn(10)
		keySequence.Keys = append(keySequence.Keys, key)
	}

	return keySequence
}
