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
	UpdateEmployee(c context.Context, out entities.Employee, id string) error
	CheckIfEmailOrPhoneNumberOrIdExists(c context.Context, email string, areaCode string, phoneNumber string, ID string, uuid string) error
	GetTitles(c context.Context) (*[]entities.Title, error)
	IsExistBasHekim(c context.Context) (bool, error)
	GetClinics(c context.Context) (*[]entities.Clinic, error)
	IsValidClinicIdBelongToHospital(c context.Context, clinicIdBelongToHospital string, hospitalId string) error
	CheckEmployeeExists(c context.Context, id string) (bool, error)
	DeleteEmployee(c context.Context, id string) error
}

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) IEmployeeRepository {
	return &EmployeeRepository{db}
}

func (er *EmployeeRepository) IsValidClinicIdBelongToHospital(c context.Context, clinicIdBelongToHospital string, hospitalId string) error {
	var clinicAndHospital entities.ClinicAndHospital
	if err := er.db.WithContext(c).Model(entities.ClinicAndHospital{}).Where("uuid = ? AND hospital_id = ?", clinicIdBelongToHospital, hospitalId).Find(&clinicAndHospital).Error; err != nil {
		return err
	}

	if clinicAndHospital.Base.UUID == uuid.Nil {
		return errors.New("clinic can't find")
	}
	return nil
}

func (er *EmployeeRepository) DeleteEmployee(c context.Context, id string) error {
	if err := er.db.WithContext(c).Where("uuid = ?", id).Delete(&entities.Employee{}).Error; err != nil {
		return err
	}
	return nil
}
func (er *EmployeeRepository) CheckEmployeeExists(c context.Context, id string) (bool, error) {
	var count int64
	if err := er.db.WithContext(c).Model(&entities.Employee{}).Where("uuid = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (er *EmployeeRepository) UpdateEmployee(c context.Context, out entities.Employee, id string) error {
	if err := er.db.WithContext(c).Model(&entities.Employee{}).Where("uuid = ?", id).Updates(&out).Error; err != nil {
		return err
	}
	return nil
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

func (er *EmployeeRepository) CheckIfEmailOrPhoneNumberOrIdExists(c context.Context, email string, areaCode string, phoneNumber string, ID string, employeeUUID string) error {
	var count int64
	var err error
	if employeeUUID == "" {
		err = er.db.WithContext(c).Model(&entities.Employee{}).Where("email = ? OR (area_code = ? AND phone = ?) OR id = ?", email, areaCode, phoneNumber, ID).Count(&count).Error
	} else {
		err = er.db.WithContext(c).Model(&entities.Employee{}).Where("email = ? OR (area_code = ? AND phone = ?) OR id = ?", email, areaCode, phoneNumber, ID).Where("uuid != ?", employeeUUID).Count(&count).Error
	}

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("email, phone number or id already exists")
	}

	return nil
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
