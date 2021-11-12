package data

import (
	"context"
	"gin-template/internal/app/data/schema"
	"gin-template/internal/app/entity"
	"strconv"
)

//go:generate mockgen -source=user.go -destination=mocks_data/user_repo_test.go -package=data_test

type UserRepo interface {
	UserLogin(ctx context.Context, userLogin *entity.UserLogin) (*entity.UserLogin, error)
	UserRegister(ctx context.Context, userRegister *entity.UserLogin) (*entity.UserLogin, error)
}

type userRepo struct {
	data *Data
}

func (u *userRepo) UserRegister(ctx context.Context, userRegister *entity.UserLogin) (*entity.UserLogin, error) {
	user := schema.User{
		UserName:    userRegister.UserName,
		Email:       userRegister.Email,
		Phone:       userRegister.Phone,
		PassWordMd5: userRegister.PassWordMd5,
	}
	err := u.data.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	userRegister.UserId = strconv.FormatInt(user.UserId, 10)
	return userRegister, err
}

func (u *userRepo) UserLogin(ctx context.Context, userLogin *entity.UserLogin) (*entity.UserLogin, error) {
	user := schema.User{
		UserName: userLogin.UserName,
		Email:    userLogin.Email,
		Phone:    userLogin.Phone,
	}
	err := u.data.db.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}
	ru := &entity.UserLogin{
		UserId:      strconv.FormatInt(user.UserId, 10),
		UserName:    user.UserName,
		Email:       user.Email,
		Phone:       user.Phone,
		PassWordMd5: user.PassWordMd5,
	}
	return ru, err
}

func NewUserRepo(data *Data) UserRepo {
	return &userRepo{
		data: data,
	}
}
