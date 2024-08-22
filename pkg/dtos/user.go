package dtos

import "github.com/yusufguntav/hospital-management/pkg/entities"

type DTOUser struct {
	ID       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	AreaCode string `json:"area_code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

type DTOUserWithRole struct {
	ID       string            `json:"id" binding:"required"`
	Email    string            `json:"email" binding:"required"`
	Phone    string            `json:"phone" binding:"required"`
	AreaCode string            `json:"area_code" binding:"required"`
	Password string            `json:"password" binding:"required"`
	Name     string            `json:"name" binding:"required"`
	Surname  string            `json:"surname" binding:"required"`
	Role     entities.AuthRole `json:"role" binding:"required"`
}

type DTOUserWithRoleAndID struct {
	UUID     string            `json:"uuid" binding:"required"`
	ID       string            `json:"id" binding:"required"`
	Email    string            `json:"email" binding:"required"`
	Phone    string            `json:"phone" binding:"required"`
	AreaCode string            `json:"area_code" binding:"required"`
	Password string            `json:"password" binding:"required"`
	Name     string            `json:"name" binding:"required"`
	Surname  string            `json:"surname" binding:"required"`
	Role     entities.AuthRole `json:"role" binding:"required"`
}

type DTOUserLogin struct {
	MailOrPhone string `json:"mail_or_phone"  binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type DTOResetPassword struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	AreaCode    string `json:"area_code" binding:"required"`
	Code        int    `json:"code" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
