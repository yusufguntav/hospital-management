package employee

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yusufguntav/hospital-management/pkg/cache"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/state"
	"gorm.io/gorm"
)

type IEmployeeRepository interface {
	Register(c context.Context, out entities.Employee) error
	CheckIfEmailOrPhoneNumberOrIdExists(c context.Context, email string, areaCode string, phoneNumber string, ID string) (entities.Employee, error)
	GetTitles(c context.Context) (*[]entities.Title, error)
	IsExistBasHekim(c context.Context) (bool, error)
	GetClinics(c context.Context) (*[]entities.Clinic, error)
	CheckClinicBelongsToHospital(c context.Context, clinicId int) (bool, error)
}

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) IEmployeeRepository {
	return &EmployeeRepository{db}
}

func (er *EmployeeRepository) CheckClinicBelongsToHospital(c context.Context, clinicId int) (bool, error) {
	var count int64
	er.db.WithContext(c).Model(entities.ClinicAndHospital{}).Where("clinic_id = ? AND hospital_id = ?", clinicId, state.CurrentUserHospitalId(c)).Count(&count)
	if count > 0 {
		return true, nil
	}
	return false, nil
}
func (er *EmployeeRepository) Register(c context.Context, out entities.Employee) error {
	if err := er.db.WithContext(c).Create(&out).Error; err != nil {
		return err
	}
	return nil
}
func (er *EmployeeRepository) IsExistBasHekim(c context.Context) (bool, error) {
	var count int64
	er.db.WithContext(c).Model(entities.Employee{}).Where("title_id = ? AND job_id = ? AND hospital_id = ?", 4, 2, state.CurrentUserHospitalId(c)).Count(&count)
	if count > 0 {
		return true, nil
	}
	return false, nil
}
func (er *EmployeeRepository) CheckIfEmailOrPhoneNumberOrIdExists(c context.Context, email string, areaCode string, phoneNumber string, ID string) (entities.Employee, error) {
	var employee entities.Employee
	er.db.WithContext(c).Model(entities.Employee{}).Where("email = ? OR (phone = ? AND area_code = ?) OR id = ?", email, phoneNumber, areaCode, ID).First(&employee)
	if employee.Base.UUID != uuid.Nil {
		return employee, errors.New("email, phone number or id already exists")
	}
	return entities.Employee{}, nil
}

func (er *EmployeeRepository) GetTitles(c context.Context) (*[]entities.Title, error) {

	cacheTitles, err := cache.GetTitles(c, er.db)
	if err != nil {
		return nil, err
	}

	return cacheTitles, nil
}

func (er *EmployeeRepository) GetClinics(c context.Context) (*[]entities.Clinic, error) {

	cacheClinics, err := cache.GetClinics(c, er.db)
	if err != nil {
		return nil, err
	}

	return cacheClinics, nil
}
