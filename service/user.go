package service

import (
	"gorm.io/gorm"
	"login-test/dao"
	"login-test/model"
)

func CheckPassword(username, password string) (error, bool) {
	err, check := dao.CheckPassword(username)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(check.Password, password)
	if err != nil {
		return err, false
	}
	return err, res
}

func CheckUsername(user model.User) (error, bool) {
	err := dao.CheckUsername(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return err, true
		}
		return err, false
	}
	return err, false
}

func WriteIn(user model.User) error {
	err := dao.WriteIn(user)
	if err != nil {
		return err
	}
	return err
}
