package user

import (
	"context"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
)

type IUserService interface {
	Register(c context.Context, req dtos.DTOUserRegister) error
}

type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(ur IUserRepository) IUserService {
	return &UserService{ur}
}

func (ur *UserService) Register(c context.Context, req dtos.DTOUserRegister) error {
	return ur.UserRepository.Register(c, req)
}
