package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golvl2/app/domain"
	"golvl2/app/domain/models"
	"golvl2/app/domain/usecases/create_user"
	"golvl2/app/domain/usecases/get_user"
)

type Handlers struct {
	r domain.Repositories
}

func NewHandlers(r domain.Repositories) *Handlers {
	return &Handlers{r: r}
}

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (h *Handlers) CreateUser(ctx context.Context, u User) ([]byte, error) {
	bu := models.User{
		Name: u.Name,
	}

	nbu, err := create_user.Run(ctx, h.r, bu)
	if err != nil {
		return nil, fmt.Errorf("error when creating: %w", err)
	}

	return json.Marshal(nbu)
}

func (h *Handlers) GetUser(ctx context.Context, id string) ([]byte, error) {

	nbu, err := get_user.Run(ctx, h.r, id)
	if err != nil {
		return nil, fmt.Errorf("error when getting: %w", err)
	}

	return json.Marshal(nbu)
}
