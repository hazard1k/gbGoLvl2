package create_user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golvl2/app/domain"
	"golvl2/app/domain/models"
)

func Run(ctx context.Context, r domain.Repositories, user models.User) (models.User, error) {

	user.ID = uuid.New()
	nbu, err := r.Users().Create(ctx, user)
	if err != nil {
		return models.User{}, fmt.Errorf("unable to create due: %w", err)
	}

	return models.User{
		ID:   nbu.ID,
		Name: nbu.Name,
	}, nil

}
