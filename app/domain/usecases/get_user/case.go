package get_user

import (
	"context"
	"fmt"
	"golvl2/app/domain"
	"golvl2/app/domain/models"
)

func Run(ctx context.Context, r domain.Repositories, id string) (models.User, error) {

	nbu, err := r.Users().Get(ctx, id)
	if err != nil {
		return models.User{}, fmt.Errorf("unable to find due: %w", err)
	}

	return models.User{
		ID:   nbu.ID,
		Name: nbu.Name,
	}, nil

}
