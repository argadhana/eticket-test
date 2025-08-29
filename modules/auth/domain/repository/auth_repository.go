package repository

import (
	"context"
	"eticket-test/modules/auth/domain/entity"
)

type AuthUserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
