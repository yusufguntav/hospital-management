package user

import (
	"context"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
)

type IUserService interface {
	Register(c context.Context, req dtos.DTOUserRegister) error
	ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error)
	ResetPassword(c context.Context, req dtos.DTOResetPassword) error
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

func (ur *UserService) ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error) {
	return ur.UserRepository.ResetPasswordApprove(c, phoneNumber, areaCode)
}

func (ur *UserService) ResetPassword(c context.Context, req dtos.DTOResetPassword) error {
	return ur.UserRepository.ResetPassword(c, req)
}
