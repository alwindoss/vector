package station

import (
	"fmt"

	"github.com/alwindoss/vector/internal/storage"
)

type Service interface {
	CreateStation(*Station) (*Station, error)
}

func NewService(repo storage.StationRepository) Service {
	return &stationService{
		Repo: repo,
	}
}

type stationService struct {
	Repo storage.StationRepository
}

// CreateStation implements Service
func (svc *stationService) CreateStation(stn *Station) (*Station, error) {
	s := &storage.Station{
		Name: stn.Name,
	}
	s, err := svc.Repo.CreateStation(s)
	if err != nil {
		err = fmt.Errorf("unable to create the record: %w", err)
		return nil, err
	}
	stn.ID = fmt.Sprintf("%d", s.ID)

	return stn, nil
}
