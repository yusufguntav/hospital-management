package hospital

import "gorm.io/gorm"

type IHospitalRepository interface {
}

type HospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) IHospitalRepository {
	return &HospitalRepository{db}
}
