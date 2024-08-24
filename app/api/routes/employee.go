package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/pkg/domains/employee"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/middleware"
)

func EmployeeRoutes(r *gin.RouterGroup, e employee.IEmployeeService) {
	r.POST("/", middleware.CheckAuth(entities.Manager, entities.Owner), employeeRegister(e))
	r.PUT("/", middleware.CheckAuth(entities.Manager, entities.Owner), employeeUpdate(e))
	r.DELETE("/:id", middleware.CheckAuth(entities.Manager, entities.Owner), employeeDelete(e))
}

func employeeDelete(e employee.IEmployeeService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := e.Delete(c, id); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Employee deleted successfully"})
	}
}

func employeeUpdate(e employee.IEmployeeService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOEmployeeWithId
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := e.Update(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Employee updated successfully"})
	}
}

func employeeRegister(e employee.IEmployeeService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOEmployee
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := e.Register(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Employee registered"})
	}
}
