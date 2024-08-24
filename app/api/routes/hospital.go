package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/pkg/domains/hospital"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/middleware"
)

func HospitalRoutes(r *gin.RouterGroup, h hospital.IHospitalService) {
	r.POST("/register", hospitalRegister(h))
	r.POST("/clinic", middleware.CheckAuth(entities.Manager, entities.Owner), addClinic(h))
	r.GET("/clinics", middleware.CheckAuth(), getClinics(h))
}

func getClinics(h hospital.IHospitalService) func(c *gin.Context) {
	return func(c *gin.Context) {
		clinicsAndEmployee, totalCount, err := h.GetClinics(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"totalCount": totalCount, "clinics": clinicsAndEmployee})
	}
}
func hospitalRegister(h hospital.IHospitalService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOHospitalRegister
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := h.Register(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Hospital registered"})
	}
}

func addClinic(h hospital.IHospitalService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req dtos.DTOClinicAdd
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := h.AddClinic(c, req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Clinic added successfully"})
	}
}
