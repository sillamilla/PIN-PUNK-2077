package service

import (
	"MiniGame-PinUp/Hacking_Service/internal/helper"
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"MiniGame-PinUp/Hacking_Service/internal/repository"
	"time"
)

type Matrix interface {
	CountMatches(matrix models.Matrix, keySequence models.KeySequence) (string, int)
	SaveHackAttempt(keys string, number int) error

	GetAll() ([]models.HackData, error)
}

type service struct {
	repository repository.Matrix
}

func NewService(repo repository.Matrix) Matrix {
	return &service{repository: repo}
}

func (s *service) CountMatches(matrix models.Matrix, keySequence models.KeySequence) (string, int) {
	HackedKey = []int{}

	horizontal(matrix, keySequence.Keys, 0, 0, false, false)

	keys := helper.SliceToString(HackedKey)

	switch {
	case len(HackedKey) >= 2 && len(HackedKey) <= 3:
		return keys, 1
	case len(HackedKey) >= 4 && len(HackedKey) <= 6:
		return keys, 2
	case len(HackedKey) == 7:
		return keys, 3
	}

	return keys, 0
}

func (s *service) GetAll() ([]models.HackData, error) {
	data, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) SaveHackAttempt(keys string, number int) error {
	hackData := models.HackData{
		Key: keys,
		Data: models.JSONData{
			Created:      time.Now(),
			ResultOfHack: number,
		},
	}

	err := s.repository.Save(hackData)
	if err != nil {
		return err
	}

	return nil
}

var HackedKey []int // Глобальная переменная для хранения найденных ключей

func horizontal(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool) {
	currentX := X
	currentY := Y

	// Двигаемся вправо (по оси Y)
	for currentY = Y; currentY < len(matrix.Data[X]); currentY++ {
		if currentY >= len(matrix.Data[currentX]) {
			break
		}

		for _, key := range sequence {
			if len(HackedKey) == 7 {
				return
			}

			if key == matrix.Data[currentX][currentY] {
				HackedKey = append(HackedKey, matrix.Data[currentX][currentY])
				matrix.Data[currentX][currentY] = -1

				if currentY == 4 {
					negativeY = true
				}
				if currentY == 0 {
					negativeY = false
				}

				if negativeX {
					verticalReverse(matrix, sequence, currentX-1, currentY, negativeX, negativeY)
					return
				} else {
					vertical(matrix, sequence, currentX+1, currentY, negativeX, negativeY)
					return
				}
			}
		}
	}
}

// Функция для обратного горизонтального движения
func horizontalReverse(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool) {
	currentX := X
	currentY := Y

	for currentY = Y; currentY >= 0; currentY-- {
		if currentY < 0 {
			break
		}

		for _, key := range sequence {
			if len(HackedKey) == 7 {
				return
			}

			if key == matrix.Data[currentX][currentY] {
				HackedKey = append(HackedKey, matrix.Data[currentX][currentY])
				matrix.Data[currentX][currentY] = -1

				if currentY == 4 {
					negativeY = true
				}
				if currentY == 0 {
					negativeY = false
				}

				if negativeX {
					verticalReverse(matrix, sequence, currentX-1, currentY, negativeX, negativeY)
					return
				} else {
					horizontal(matrix, sequence, currentX+1, currentY, negativeX, negativeY)
					return
				}
			}
		}
	}
}

func vertical(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool) {
	currentX := X
	currentY := Y

	// Двигаемся вниз (по оси X)
	for currentX = X; currentX < len(matrix.Data); currentX++ {
		if currentX >= len(matrix.Data[currentY]) {
			break
		}

		for _, key := range sequence {
			if len(HackedKey) == 7 {
				return
			}

			if key == matrix.Data[currentX][Y] {
				HackedKey = append(HackedKey, matrix.Data[currentX][Y])
				matrix.Data[currentX][currentY] = -1

				if currentX == 4 {
					negativeX = true
				}
				if currentX == 0 {
					negativeX = false
				}

				if negativeY {
					horizontalReverse(matrix, sequence, currentX, currentY-1, negativeX, negativeY)
					return
				} else {
					horizontal(matrix, sequence, currentX, currentY+1, negativeX, negativeY)
					return
				}
			}
		}
	}
}

// Функция для обратного вертикального движения
func verticalReverse(matrix models.Matrix, sequence []int, X, Y int, negativeX, negativeY bool) {
	currentX := X
	currentY := Y

	for currentX = X; currentX >= 0; currentX-- {
		if currentX < 0 {
			break
		}

		for _, key := range sequence {
			if len(HackedKey) == 7 {
				return
			}

			if key == matrix.Data[currentX][currentY] {
				HackedKey = append(HackedKey, matrix.Data[currentX][currentY])
				matrix.Data[currentX][currentY] = -1

				if currentX == 0 {
					negativeX = false
				}
				if currentX == 4 {
					negativeX = false
				}

				if negativeY {
					horizontalReverse(matrix, sequence, currentX, currentY-1, negativeX, negativeY)
					return
				} else {
					horizontal(matrix, sequence, currentX, currentY+1, negativeX, negativeY)
					return
				}
			}
		}
	}
}
