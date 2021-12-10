package data

import (
	"context"
	"gin-template/internal/app/data/schema"
	"gin-template/pkg/logging"
	"strconv"
)

func UserRegister(ctx context.Context, user schema.User) (string, error) {
	user.GenerateId()
	logging.Info(user.UserId)
	err := GetDB().Save(&user).Error
	if err != nil {
		return "", err
	}
	id := strconv.FormatInt(user.UserId, 10)
	return id, err
}
