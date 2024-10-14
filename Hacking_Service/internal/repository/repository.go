package repository

import (
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type Matrix interface {
	Save(hackResult models.HackData) error
	GetAll() ([]models.HackData, error)
}

type repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) Matrix {
	return &repository{db: db}
}

func (r *repository) Save(hackResult models.HackData) error {
	_, err := r.db.Model(&hackResult).Insert()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *repository) GetAll() ([]models.HackData, error) {
	var hackDataList []models.HackData

	err := r.db.Model(&hackDataList).Select()

	if err != nil {
		return nil, err
	}

	return hackDataList, nil
}
