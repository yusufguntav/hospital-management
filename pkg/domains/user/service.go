package user

import (
	"context"
	"errors"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/state"
)

type IUserService interface {
	RegisterSubUser(c context.Context, req dtos.DTOUserWithRole) error
	Login(c context.Context, req dtos.DTOUserLogin) (string, error)
	ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error)
	ResetPassword(c context.Context, req dtos.DTOResetPassword) error
	UpdateUser(c context.Context, req dtos.DTOUserWithRoleAndID) error
	DeleteSubUser(c context.Context, id string) error
}

type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(ur IUserRepository) IUserService {
	return &UserService{ur}
}
func (ur *UserService) DeleteSubUser(c context.Context, id string) error {

	// Check if subuser exist
	user, err := ur.UserRepository.CheckIfUserExists(c, id)

	if err != nil {
		return err
	}

	// Check authorization
	if user.Role == entities.Owner {
		return errors.New("can't delete owner")
	}

	if user.Role == entities.Manager && state.CurrentUserRole(c) != entities.Owner {
		return errors.New("you can't delete manager")
	}

	// Delete subuser
	if err := ur.UserRepository.DeleteSubUser(c, id); err != nil {
		return err
	}

	return nil
}

func (ur *UserService) UpdateUser(c context.Context, req dtos.DTOUserWithRoleAndID) error {
	return ur.UserRepository.UpdateUser(c, req)
}
func (ur *UserService) Login(c context.Context, req dtos.DTOUserLogin) (string, error) {
	return ur.UserRepository.Login(c, req)
}

func (ur *UserService) RegisterSubUser(c context.Context, req dtos.DTOUserWithRole) error {
	return ur.UserRepository.RegisterSubUser(c, req)
}

func (ur *UserService) ResetPasswordApprove(c context.Context, phoneNumber string, areaCode string) (int, error) {
	return ur.UserRepository.ResetPasswordApprove(c, phoneNumber, areaCode)
}

func (ur *UserService) ResetPassword(c context.Context, req dtos.DTOResetPassword) error {
	return ur.UserRepository.ResetPassword(c, req)
}
