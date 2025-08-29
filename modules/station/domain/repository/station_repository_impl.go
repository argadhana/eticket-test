package repository

import (
	"context"
	"eticket-test/modules/station/domain/entity"
	"gorm.io/gorm"
)

type stationRepositoryImpl struct {
	db *gorm.DB
}

func NewStationRepositoryImpl(db *gorm.DB) StationRepository {
	return &stationRepositoryImpl{db: db}
}

func (r *stationRepositoryImpl) Create(ctx context.Context, station *entity.Station) error {
	return r.db.WithContext(ctx).Create(station).Error
}
