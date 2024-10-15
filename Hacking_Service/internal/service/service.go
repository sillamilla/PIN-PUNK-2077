package service

import (
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"MiniGame-PinUp/Hacking_Service/internal/repository"
	"context"
	"strconv"
	"strings"
	"time"
)

type MatrixService interface {
	Hack(ctx context.Context, matrix models.Matrix, keySequence models.KeySequence) (string, error)
	SaveHack(ctx context.Context, key []int, resultOfHack int) error

	GetAll(ctx context.Context) ([]models.HackData, error)
}

type service struct {
	repository repository.MatrixRepository
}

func New(repo repository.MatrixRepository) MatrixService {
	return &service{repository: repo}
}

func (s *service) Hack(ctx context.Context, matrix models.Matrix, keySequence models.KeySequence) (string, error) {
	key, resultOfHack := s.hack(matrix, keySequence)

	status := "fail"
	if resultOfHack > 0 {
		status = "success"

		if err := s.SaveHack(ctx, key, resultOfHack); err != nil {
			return "", err
		}
	}

	return status, nil
}

func (s *service) SaveHack(ctx context.Context, key []int, resultOfHack int) error {
	hackData := models.HackData{
		Key: sliceToString(key),
		Data: models.JSONData{
			Created:      time.Now(),
			ResultOfHack: resultOfHack,
		},
	}

	if err := s.repository.Save(ctx, hackData); err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context) ([]models.HackData, error) {
	if data, err := s.repository.GetAll(ctx); err == nil {
		return data, nil
	} else {
		return nil, err
	}
}

func (s *service) hack(matrix models.Matrix, keySequence models.KeySequence) ([]int, int) {
	hackedKey := make([]int, 0, models.KeySequenceLength)

	s.horizontal(matrix, keySequence.Keys, 0, 0, false, false, &hackedKey)

	resultOfHack := 0
	switch len(hackedKey) {
	case 2, 3:
		resultOfHack = 1
	case 4, 5, 6:
		resultOfHack = 2
	case 7:
		resultOfHack = 3
	}

	return hackedKey, resultOfHack
}

func (s *service) horizontal(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool, hackedKey *[]int) {
	currentX := X
	currentY := Y

	// Двигаемся вправо (по оси Y)
	for currentY = Y; currentY < len(matrix.Data[X]); currentY++ {
		if currentY >= len(matrix.Data[currentX]) {
			break
		}

		for _, key := range sequence {
			if len(*hackedKey) == models.KeySequenceLength {
				return
			}

			if key == matrix.Data[currentX][currentY] {
				*hackedKey = append(*hackedKey, matrix.Data[currentX][currentY])
				matrix.Data[currentX][currentY] = -1

				if currentY == 4 {
					negativeY = true
				}
				if currentY == 0 {
					negativeY = false
				}

				if negativeX {
					s.verticalReverse(matrix, sequence, currentX-1, currentY, negativeX, negativeY, hackedKey)
					return
				} else {
					s.vertical(matrix, sequence, currentX+1, currentY, negativeX, negativeY, hackedKey)
					return
				}
			}
		}
	}
}

// Функция для обратного горизонтального движения
func (s *service) horizontalReverse(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool, hackedKey *[]int) {
	currentX := X
	currentY := Y

	for currentY = Y; currentY >= 0; currentY-- {
		if currentY < 0 {
			break
		}

		for _, key := range sequence {
			if len(*hackedKey) == models.KeySequenceLength {
				return
			}

			if key == matrix.Data[currentX][currentY] {
				*hackedKey = append(*hackedKey, matrix.Data[currentX][currentY])
				matrix.Data[currentX][currentY] = -1

				if currentY == 4 {
					negativeY = true
				}
				if currentY == 0 {
					negativeY = false
				}

				if negativeX {
					s.verticalReverse(matrix, sequence, currentX-1, currentY, negativeX, negativeY, hackedKey)
					return
				} else {
					s.horizontal(matrix, sequence, currentX+1, currentY, negativeX, negativeY, hackedKey)
					return
				}
			}
		}
	}
}

func (s *service) vertical(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool, hackedKey *[]int) {
	currentX := X
	currentY := Y

	// Двигаемся вниз (по оси X)
	for currentX = X; currentX < len(matrix.Data); currentX++ {
		if currentX >= len(matrix.Data[currentY]) {
			break
		}

		for _, key := range sequence {
			if len(*hackedKey) == models.KeySequenceLength {
				return
			}

			if key == matrix.Data[currentX][Y] {
				*hackedKey = append(*hackedKey, matrix.Data[currentX][Y])
				matrix.Data[currentX][currentY] = -1

				if currentX == 4 {
					negativeX = true
				}
				if currentX == 0 {
					negativeX = false
				}

				if negativeY {
					s.horizontalReverse(matrix, sequence, currentX, currentY-1, negativeX, negativeY, hackedKey)
					return
				} else {
					s.horizontal(matrix, sequence, currentX, currentY+1, negativeX, negativeY, hackedKey)
					return
				}
			}
		}
	}
}

// Функция для обратного вертикального движения
func (s *service) verticalReverse(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool, hackedKey *[]int) {
	currentX := X
	currentY := Y

	for currentX = X; currentX >= 0; currentX-- {
		if currentX < 0 {
			break
		}

		for _, key := range sequence {
			if len(*hackedKey) == models.KeySequenceLength {
				return
			}

			if key == matrix.Data[currentX][currentY] {
				*hackedKey = append(*hackedKey, matrix.Data[currentX][currentY])
				matrix.Data[currentX][currentY] = -1

				if currentX == 0 {
					negativeX = false
				}
				if currentX == 4 {
					negativeX = false
				}

				if negativeY {
					s.horizontalReverse(matrix, sequence, currentX, currentY-1, negativeX, negativeY, hackedKey)
					return
				} else {
					s.horizontal(matrix, sequence, currentX, currentY+1, negativeX, negativeY, hackedKey)
					return
				}
			}
		}
	}
}

func sliceToString(slice []int) string {
	var builder strings.Builder
	for _, num := range slice {
		builder.WriteString(strconv.Itoa(num))
	}
	return builder.String()
}
