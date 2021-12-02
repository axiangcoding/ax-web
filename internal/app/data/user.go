package data

import (
	"context"
	"gin-template/internal/app/data/schema"
	"strconv"
)

func UserRegister(ctx context.Context) (string, error) {
	email := "test"
	phone := "phone"
	user := schema.User{
		UserName: "Abc",
		Email:    &email,
		Phone:    &phone,
		Password: "abcdedf",
	}
	err := GetDB().Save(&user).Error
	if err != nil {
		return "", err
	}
	id := strconv.FormatInt(int64(user.UserId), 10)
	return id, err
}
