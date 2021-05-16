package domain

import "context"

type User struct {
	ID   int
	Name string
	Age int64
}

type IUserUsecase interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}

type IUserRepo interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}
