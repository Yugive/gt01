package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"week2/internal/errcode"
	"week2/internal/model"
)

func GetUserById(id int) (model.User, error) {
	user := model.User{
		ID: id,
	}
	user, err := user.GetUser(model.UserDB)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.Wrapf(errcode.UserNotExist, "record not found: %+v", err)
		} else {
			return user, errors.Wrapf(errcode.DBQuery, "db errors is: %+v", err)
		}
	}

	return user, nil
}
