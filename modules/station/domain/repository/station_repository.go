package repository

import (
	"context"
	"eticket-test/modules/station/domain/entity"
)

type StationRepository interface {
	Create(ctx context.Context, station *entity.Station) error
}
