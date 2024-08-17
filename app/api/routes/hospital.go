package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/pkg/domains/hospital"
)

func HospitalRoutes(r *gin.RouterGroup, s hospital.IHospitalService) {
	r.GET("/x", x(s))
}

func x(s hospital.IHospitalService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "x"})
	}
}
