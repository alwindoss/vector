package storage

import (
	"fmt"

	"gorm.io/gorm"
)

type StationRepository interface {
	CreateStation(s *Station) (*Station, error)
}

func NewStationRepository(db *gorm.DB) StationRepository {
	return &vectorStationRepository{
		DB: db,
	}
}

type vectorStationRepository struct {
	DB *gorm.DB
}

// CreateStation implements StationRepository
func (sr *vectorStationRepository) CreateStation(s *Station) (*Station, error) {
	tx := sr.DB.Create(s)
	if tx.Error != nil {
		err := fmt.Errorf("unable to create the record in the DB: %w", tx.Error)
		return nil, err
	}
	return s, nil
}
