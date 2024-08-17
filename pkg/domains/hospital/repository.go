package hospital

import (
	"context"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IHospitalRepository interface {
	Register(c context.Context, req dtos.DTOHospitalRegister) error
}

type HospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) IHospitalRepository {
	return &HospitalRepository{db}
}

func (ur *HospitalRepository) Register(c context.Context, req dtos.DTOHospitalRegister) error {
	// Password hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Manager.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Start transaction
	tx := ur.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Owner user creation
	entUser := entities.User{
		ID:      req.Manager.ID,
		Name:    req.Manager.Name,
		Surname: req.Manager.Surname,
		Contact: entities.Contact{Email: req.Manager.Email, Phone: req.Manager.Phone, AreaCode: req.Manager.AreaCode},
		Role:    entities.Owner,
	}
	entUser.Password = string(passwordHash)
	if err := tx.WithContext(c).Create(&entUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hospital creation
	entHospital := entities.Hospital{
		TID:          req.HTID,
		Name:         req.HName,
		Address:      req.HAddress,
		CityCode:     req.HCityCode,
		DistrictCode: req.HDistrictCode,
		ManagerId:    entUser.Base.UUID.String(),
		Contact:      entities.Contact{Email: req.HEmail, Phone: req.HPhone, AreaCode: req.HAreaCode},
	}

	if err := tx.WithContext(c).Create(&entHospital).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
