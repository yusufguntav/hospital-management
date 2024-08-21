package dtos

import "github.com/yusufguntav/hospital-management/pkg/entities"

type DTOOwnerRegister struct {
	ID       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	AreaCode string `json:"area_code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

type DTOUserRegister struct {
	ID       string            `json:"id" binding:"required"`
	Email    string            `json:"email" binding:"required"`
	Phone    string            `json:"phone" binding:"required"`
	AreaCode string            `json:"area_code" binding:"required"`
	Password string            `json:"password" binding:"required"`
	Name     string            `json:"name" binding:"required"`
	Surname  string            `json:"surname" binding:"required"`
	Role     entities.AuthRole `json:"role" binding:"required"`
}
