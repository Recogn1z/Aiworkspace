package dao

import (
	"ai-workspace-backend/common"
	"ai-workspace-backend/model"

	"gorm.io/gorm"
)

func CreateUser(user *model.User) error {
	return common.DB.Create(user).Error
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := common.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsRecordNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}