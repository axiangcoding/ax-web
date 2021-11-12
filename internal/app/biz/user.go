package biz

import (
	"context"
	"gin-template/internal/app/data"
	"gin-template/internal/app/entity"
)

type UserUseCase interface {
	UserLogin(ctx context.Context, login *entity.UserLogin) (*entity.UserLogin, error)
	UserRegister(ctx context.Context, login *entity.UserLogin) (*entity.UserLogin, error)
}

type userUseCase struct {
	repo data.UserRepo
}

func (u *userUseCase) UserRegister(ctx context.Context, login *entity.UserLogin) (*entity.UserLogin, error) {
	return u.repo.UserRegister(ctx, login)
}

func (u *userUseCase) UserLogin(ctx context.Context, login *entity.UserLogin) (*entity.UserLogin, error) {
	return u.repo.UserLogin(ctx, login)
}

func NewUserUseCase(repo data.UserRepo) UserUseCase {
	return &userUseCase{repo: repo}
}
