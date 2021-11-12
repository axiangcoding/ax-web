package biz

import (
	"gin-template/internal/app/data"
)

type UserUseCase interface {
}

type userUseCase struct {
	repo data.UserRepo
}

func NewUserUseCase(repo data.UserRepo) UserUseCase {
	return &userUseCase{repo: repo}
}
