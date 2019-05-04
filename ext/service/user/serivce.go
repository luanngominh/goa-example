package user

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"github.com/luanngominh/goa-example/model"
)

type Service interface {
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Get(ctx context.Context, q *UserQuery) (*model.User, error)
}

type UserQuery struct {
	Email string
	ID uuid.UUID
}