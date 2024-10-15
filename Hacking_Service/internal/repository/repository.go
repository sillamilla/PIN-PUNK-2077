package repository

import (
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type MatrixRepository interface {
	Save(ctx context.Context, hackResult models.HackData) error
	GetAll(ctx context.Context) ([]models.HackData, error)
}

type repository struct {
	db *pg.DB
}

func New(db *pg.DB) MatrixRepository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, hackResult models.HackData) error {
	if _, err := r.db.ModelContext(ctx, &hackResult).Insert(); err != nil {
		return fmt.Errorf("error while saving hack result: %w", err)
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]models.HackData, error) {
	var hackDataList []models.HackData
	if err := r.db.ModelContext(ctx, &hackDataList).Select(); err != nil {
		return nil, fmt.Errorf("error while getting all hack results: %w", err)
	}

	return hackDataList, nil
}
