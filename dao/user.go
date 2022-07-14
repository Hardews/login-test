package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"login-test/model"
)

func CheckPassword(username string) (error, model.User) {
	var check model.User
	tx := dB.Where("username = ?", username).First(&model.User{}).Scan(&check)
	if err := tx.Error; err != nil {
		return err, check
	}
	return nil, check
}

func CheckUsername(user model.User) error {
	var uUser model.User
	tx := dB.Where("username = ?", user.Username).First(&model.User{}).Scan(&uUser)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func WriteIn(user model.User) error {
	tx := dB.Begin()

	dx := tx.Create(&user)
	if err := dx.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
