package service

import (
	"context"
	"eticket-test/modules/station/domain/entity"
	"eticket-test/modules/station/domain/repository"
)

type StationService struct {
	stationRepo repository.StationRepository
}

func NewStationService(stationRepo repository.StationRepository) *StationService {
	return &StationService{stationRepo: stationRepo}
}

func (s *StationService) CreateStation(ctx context.Context, name, location string) (*entity.Station, error) {
	station := &entity.Station{
		Name:     name,
		Location: location,
	}
	err := s.stationRepo.Create(ctx, station)
	if err != nil {
		return nil, err
	}
	return station, nil
}
