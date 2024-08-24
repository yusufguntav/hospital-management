package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/pkg/domains/user"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/middleware"
)

func UserRoutes(r *gin.RouterGroup, u user.IUserService) {
	r.POST("/sub-user/register", middleware.CheckAuth(entities.Owner, entities.Manager), subUserRegister(u))
	r.POST("/login", userLogin(u))
	r.POST("/password-approve/:areaCode/:phoneNumber", userResetPasswordApprove(u))
	r.POST("/password-reset", userResetPassword(u))
	r.PUT("/update", middleware.CheckAuth(entities.Owner, entities.Manager), userUpdate(u))
	r.DELETE("/sub-user/:id", middleware.CheckAuth(entities.Owner, entities.Manager), userDelete(u))
}

func userDelete(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")

		if err := u.DeleteSubUser(c, id); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "User deleted"})
	}
}

func userLogin(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOUserLogin
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		token, err := u.Login(c, req)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"token": token})
	}
}
func subUserRegister(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOUserWithRole
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := u.RegisterSubUser(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "User registered"})
	}
}

func userUpdate(u user.IUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOUserWithRoleAndID
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := u.UpdateUser(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User updated"})
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
