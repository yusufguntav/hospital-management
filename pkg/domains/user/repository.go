package user

import (
	"context"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(c context.Context, req dtos.DTOUserRegister) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Register(c context.Context, req dtos.DTOUserRegister) error {
	return nil
}
