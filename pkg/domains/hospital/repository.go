package hospital

import (
	"context"
	"errors"

	"github.com/yusufguntav/hospital-management/pkg/cache"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IHospitalRepository interface {
	Register(c context.Context, req dtos.DTOHospitalRegister) error
	AddClinic(c context.Context, clinicId int, hospitalId string) error
	GetClinics(c context.Context) (*[]entities.Clinic, error)
	GetClinicsBelongingToTheHospital(c context.Context, hospitalId string) (*[]dtos.DTOClinicBelongToHospital, error)
	GetCountOfEmployeesOfEachClinic(c context.Context, clinics *[]dtos.DTOClinicBelongToHospital) (*[]dtos.DTOClinics, error)
	IsClinicAlreadyAdded(c context.Context, clinicId int, hospitalId string) (bool, error)
	GetTotalCountOfEmployees(c context.Context, hospitalId string) (int64, error)
}

type HospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) IHospitalRepository {
	return &HospitalRepository{db}
}

func (ur *HospitalRepository) GetTotalCountOfEmployees(c context.Context, hospitalId string) (int64, error) {
	var count int64
	if err := ur.db.WithContext(c).Model(&entities.Employee{}).Where("hospital_id = ?", hospitalId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *HospitalRepository) GetCountOfEmployeesOfEachClinic(c context.Context, clinics *[]dtos.DTOClinicBelongToHospital) (*[]dtos.DTOClinics, error) {
	var allJobAndEmployees []dtos.DTOClinics

	for _, clinic := range *clinics {
		var jobAndEmployees []dtos.DTOJobAndEmployee
		if err := ur.db.WithContext(c).Raw(`
        SELECT job.name as job_name, count(*) as employee_count
        FROM employee
        JOIN job on employee.job_id = job.id
        WHERE employee.deleted_at is NULL AND employee.clinic_id = ?
        GROUP BY job.name`, clinic.UUID).Find(&jobAndEmployees).Error; err != nil {
			return nil, err
		}
		allJobAndEmployees = append(allJobAndEmployees, dtos.DTOClinics{ClinicName: clinic.Name, JobAndEmployee: jobAndEmployees})
	}

	return &allJobAndEmployees, nil
}

func (ur *HospitalRepository) GetClinicsBelongingToTheHospital(c context.Context, hospitalId string) (*[]dtos.DTOClinicBelongToHospital, error) {
	var clinics []dtos.DTOClinicBelongToHospital
	if err := ur.db.WithContext(c).Raw(`
	SELECT cah.uuid,clinic.name
	FROM clinic_and_hospitals as cah
	JOIN clinic on cah.clinic_id = clinic.id
	WHERE cah.deleted_at is NULL AND cah.hospital_id = ?`, hospitalId).Find(&clinics).Error; err != nil {
		return nil, err
	}
	return &clinics, nil
}

func (ur *HospitalRepository) IsClinicAlreadyAdded(c context.Context, clinicId int, hospitalId string) (bool, error) {
	var count int64
	if err := ur.db.WithContext(c).Model(&entities.ClinicAndHospital{}).Where("clinic_id = ? AND hospital_id = ?", clinicId, hospitalId).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
func (ur *HospitalRepository) AddClinic(c context.Context, clinicId int, hospitalId string) error {
	if err := ur.db.WithContext(c).Create(&entities.ClinicAndHospital{ClinicId: clinicId, HospitalId: hospitalId}).Error; err != nil {
		return err
	}

	return nil
}

func (ur *HospitalRepository) GetClinics(c context.Context) (*[]entities.Clinic, error) {

	cacheClinics, err := cache.GetClinics(c, ur.db)
	if err != nil {
		return nil, err
	}

	return cacheClinics, nil
}
func (ur *HospitalRepository) Register(c context.Context, req dtos.DTOHospitalRegister) error {

	cacheDistricts, _, err := cache.GetDistrictsAndCities(c, ur.db)

	if err != nil {
		return err
	}

	// Check if district code is valid
	isCityAndDistrictValid := false
	for _, district := range *cacheDistricts {
		if district.ID == req.HDistrictCode && district.CityId == req.HCityCode {
			isCityAndDistrictValid = true
			break
		}
	}

	if !isCityAndDistrictValid {
		return errors.New("invalid city or district code")
	}

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

	// Hospital creation
	entHospital := entities.Hospital{
		TID:          req.HTID,
		Name:         req.HName,
		Address:      req.HAddress,
		CityCode:     req.HCityCode,
		DistrictCode: req.HDistrictCode,
		Contact:      entities.Contact{Email: req.HEmail, Phone: req.HPhone, AreaCode: req.HAreaCode},
	}

	if err := tx.WithContext(c).Create(&entHospital).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Owner user creation
	entUser := entities.User{
		ID:         req.Manager.ID,
		Name:       req.Manager.Name,
		Surname:    req.Manager.Surname,
		Contact:    entities.Contact{Email: req.Manager.Email, Phone: req.Manager.Phone, AreaCode: req.Manager.AreaCode},
		Role:       entities.Owner,
		HospitalId: entHospital.Base.UUID.String(),
	}

	entUser.Password = string(passwordHash)
	if err := tx.WithContext(c).Create(&entUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
