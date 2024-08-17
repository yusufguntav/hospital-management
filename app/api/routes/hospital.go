package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/pkg/domains/hospital"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
)

func HospitalRoutes(r *gin.RouterGroup, s hospital.IHospitalService) {
	r.POST("/register", register(s))
}

func register(s hospital.IHospitalService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOHospitalRegister
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := s.Register(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Hospital registered"})
	}
}
