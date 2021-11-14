package memory

import (
	"context"
	"fmt"
	"golvl2/app/domain/models"
	"sync"
)

type usersRepository struct {
	users *sync.Map
}

func (u *usersRepository) Get(ctx context.Context, id string) (models.User, error) {
	us, ok := u.users.Load(id)
	if !ok {
		return models.User{}, fmt.Errorf("not found")
	}

	return us.(models.User), nil
}

func (u *usersRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	u.users.Store(user.ID.String(), user)
	return user, nil
}
