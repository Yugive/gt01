package usecase

import (
	"context"

	"week4/internal/domain"
)

type user struct {
	repo domain.IUserRepo
}


func NewUserUsecase(repo domain.IUserRepo) domain.IUserUsecase {
	return &user{repo: repo}
}

func (u *user) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetUserInfo(ctx, id)
}
