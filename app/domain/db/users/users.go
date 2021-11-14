package users

import (
	"context"
	"golvl2/app/domain/models"
)

type Repository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	Get(ctx context.Context, id string) (models.User, error)
}
