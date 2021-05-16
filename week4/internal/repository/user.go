package repository

import (
	"context"
	er "week4/errors"
	"week4/internal/domain"
	"week4/internal/repository/ent"

	"github.com/pkg/errors"
)

type repository struct {
	client *ent.Client
}

// NewRepository NewRepository
func NewRepository(client *ent.Client) domain.IUserRepo {
	return &repository{client: client}
}

// GetUserInfo GetUserInfo
func (r *repository) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	user, err := r.client.User.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, errors.Wrapf(er.UserNotFound, "user not found, id: %d, err: %+v", id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(er.ErrUnKnown, "db query err: %+v, id: %d,", err, id)
	}

	return &domain.User{
		Name: user.Name,
		Age: user.Age,
		ID:   user.ID,
	}, nil
}
