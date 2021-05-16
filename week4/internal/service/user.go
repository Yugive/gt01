package service

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	v1 "week4/api/user/v1"
	er "week4/errors"
	"week4/internal/domain"
)

type UserService struct {
	v1.UserServerServer
	usecase domain.IUserUsecase
}

func NewUserService(usecase domain.IUserUsecase) *UserService {
	return &UserService{usecase: usecase}
}

func (u *UserService) GetUserInfo(ctx context.Context, req *v1.GetUerInfoRequest) (*v1.GetUserInfoReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.Wrap(er.ErrUnKnown, "get metadata err")
	}

	data := md.Get("uid")
	if len(data) != 1 {
		return nil, errors.Wrapf(er.ErrUnKnown, "user id lens not 1, metadata: %v", data)
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.Wrapf(er.ErrUnKnown, "user id not a num, data: %v", data)
	}

	user, err := u.usecase.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &v1.GetUserInfoReply{
		Username: user.Name,
		Age:     user.Age,
	}

	return resp, nil
}
