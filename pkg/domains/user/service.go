package user

import (
	"context"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
)

type IUserService interface {
	RegisterSubUser(c context.Context, req dtos.DTOSubUserRegister) error
	Login(c context.Context, req dtos.DTOUserLogin) (string, error)
	ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error)
	ResetPassword(c context.Context, req dtos.DTOResetPassword) error
}

type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(ur IUserRepository) IUserService {
	return &UserService{ur}
}

func (ur *UserService) Login(c context.Context, req dtos.DTOUserLogin) (string, error) {
	return ur.UserRepository.Login(c, req)
}

func (ur *UserService) RegisterSubUser(c context.Context, req dtos.DTOSubUserRegister) error {
	return ur.UserRepository.RegisterSubUser(c, req)
}

func (ur *UserService) ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error) {
	return ur.UserRepository.ResetPasswordApprove(c, phoneNumber, areaCode)
}

func (ur *UserService) ResetPassword(c context.Context, req dtos.DTOResetPassword) error {
	return ur.UserRepository.ResetPassword(c, req)
}
