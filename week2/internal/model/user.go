package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func (u User) TableName() string {
	return "users"
}

func (u User) GetUser(db *gorm.DB) (User, error) {
	var user User
	db = UserDB.Where("id = ?", u.ID)
	err := db.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
