package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/pkg/domains/user"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/middleware"
)

func UserRoutes(r *gin.RouterGroup, u user.IUserService) {
	r.POST("/register", userRegister(u), middleware.CheckAuth(entities.Owner, entities.Manager))
	r.POST("/password-approve/:areaCode/:phoneNumber", userResetPasswordApprove(u))
	r.POST("/password-reset", userResetPassword(u))
}

func userRegister(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOUserRegister
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := u.Register(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "User registered"})
	}
}

func userResetPasswordApprove(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {

		phoneNumber := c.Param("phoneNumber")
		areaCode := c.Param("areaCode")

		code, err := u.ResetPasswordApprove(c, phoneNumber, areaCode)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"code": code})
	}
}

func userResetPassword(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {

		var req dtos.DTOResetPassword
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := u.ResetPassword(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"code": "Password reset successfully"})
	}
}
