package repository

import (
	"github.com/rifkir23/MjTest/config"
	"github.com/rifkir23/MjTest/entity"
)

var db = config.SetupDatabaseConnection()

type UserRepository interface {
	VerifyCredential(email string) interface{}
}

type userRepo struct {
}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (u userRepo) VerifyCredential(email string) interface{} {
	var user entity.User
	res := db.Where("user_name = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
